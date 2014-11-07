/*
 Based on the generic Arduino i2c slave, this sketch extends it to handling various weather sensors
 as used in the actual Mark II project.
 */
#include <Wire.h>

// The slave address
#define SLAVE_ADDRESS 0x04

int output_size;
byte output[4];

// Debounce period
#define DEBOUNCE 500

// Anenometer pin & interrupt
#define ANEMOMETER_PIN 2
#define ANEMOMETER_INT 0

// Anenometer counts per second to speed
#define WIND_MPH 1.492
#define WIND_KPH 2.4

// Rain gauge measures 0.2794mm per count
#define RAIN_INCH 0.011
#define RAIN_MM   0.2794

// The sample time in seconds, usually we read once every 60 seconds
#define SAMPLE_TIME 60

// ==================================================
// I2C command codes
#define CMD_CPU_TEMP   0x02
#define CMD_WIND_DIR   0x20

/*
  The sensors, here each command is a block of 4 bits, lower 4 bits determine the output.
  
  So for bits 0-3:
    0  The raw measurement
    1  The measurement in metric, e.g. km/h or mm
    2  The measurement in imperial, e.g. mph or inches
    3  The measurement without any conversion
  
  Note: for 0 the response is an integer while for the others it's a word double
*/
#define CMD_MASK       0xfc
#define CMD_WIND       0x24
#define CMD_WIND_GUST  0x28
#define CMD_RAIN       0x2C

// ==================================================
// Anenometer
// ==================================================
// Initial values
#define ANEM_COUNT_INIT 0
#define ANEM_MIN_INIT   0xffffffff

volatile unsigned long anem_last  = 0;
volatile unsigned long anem_count = ANEM_COUNT_INIT;
volatile unsigned long anem_min   = ANEM_MIN_INIT;
volatile unsigned long anem = 0;
volatile unsigned long anem_gust = 0;

// ==================================================
// Rain Gauge
// ==================================================
volatile unsigned long rain_count = 0;
volatile unsigned long rain_last  = 0;
volatile unsigned long rain = 0;

// ==================================================
// Used by 1Hz timer to know when within every minute
// a sample is to be taken
unsigned int timer_second=0;

/*
 * 1Hz Timer. Here we take a snapshot of the current wind/rain readings
 * once every 60 seconds and reset the counters used by the interrupts.
 */
ISR(TIMER1_COMPA_vect) {
  if(timer_second) {
    timer_second--;
  } else {
    // Mark next sample time
    timer_second=SAMPLE_TIME;
    
    // read & reset anenometer
    anem = anem_count;
    anem_gust=anem_min;
    anem_count = ANEM_COUNT_INIT;
    anem_min = ANEM_MIN_INIT;
    
    // read & reset rain
    rain=rain_count;
    rain_count=0;
  }
}

// ==================================================

void anemometerClick() {
  long time=micros()-anem_last;
  anem_last=micros();
  if(time>DEBOUNCE) {
    anem_count++;
    if(time<anem_min) {
      anem_min=time;
    }
  }
}

void windout(int command) {
  command=command & 0x3;
  if(command==0) {
      intout(anem);
  } else {
    unitout( anem/SAMPLE_TIME, command, WIND_KPH, WIND_MPH );
  }
}

void gustout(int command) {
  command=command & 0x3;
  if(command==0) {
      intout(anem_gust);
  } else {
    unitout( 1 / (anem_gust/1000000.0), command, WIND_KPH, WIND_MPH );
  }
}

// ==================================================

void rainClick() {
  long time = micros()-rain_last;
  rain_last = micros();
  if(time>500) {
    rain_count++;
  }
}

void rainout(int command) {
  command=command & 0x3;
  if(command==0) {
      intout(rain);
  }else {
    unitout( (double)rain, command, RAIN_MM, RAIN_INCH );
  }
}

/*
  Utility which returns the appropriate value depending on the command
  m=1  returns val in metric units
  m=2  returns val in imperial units
  m=3  returns val
 */
void unitout(double val,int m,double metric,double imperial) {
  if(m==1) {
    val*=metric;
  } else if(m==2) {
    val*=imperial;
  }
  doubleout(val);
}

// ==================================================

void setup() {
  
  // Setup our timer so we reset the anenometer & rain gauge every minute
  // Disable
  cli();
  //set timer1 interrupt at 1Hz
  TCCR1A = 0;// set entire TCCR1A register to 0
  TCCR1B = 0;// same for TCCR1B
  TCNT1  = 0;//initialize counter value to 0
  // set compare match register for 1hz increments
  OCR1A = 15624;// = (16*10^6) / (1*1024) - 1 (must be <65536)
  // turn on CTC mode
  TCCR1B |= (1 << WGM12);
  // Set CS10 and CS12 bits for 1024 prescaler
  TCCR1B |= (1 << CS12) | (1 << CS10);  
  // enable timer compare interrupt
  TIMSK1 |= (1 << OCIE1A);
  // Enable interrupts
  sei();
  
  // initialize i2c as slave
  Wire.begin(SLAVE_ADDRESS);
  
  // define callbacks for i2c communication
  Wire.onReceive(receiveData);
  Wire.onRequest(sendData);

  // Set onboard LED
  pinMode(LED_BUILTIN,OUTPUT);

  // Setup anenometer
  pinMode(ANEMOMETER_PIN,INPUT);
  digitalWrite(ANEMOMETER_PIN,HIGH); // turn on internal pull up resistor
  attachInterrupt(ANEMOMETER_INT, anemometerClick, FALLING);
  interrupts();
  
  // The internal temperature has to be used
  // with the internal reference of 1.1V.
  // Channel 8 can not be selected with
  // the analogRead function yet.

  // Set the internal reference and mux.
  ADMUX = (_BV(REFS1) | _BV(REFS0) | _BV(MUX3));
  ADCSRA |= _BV(ADEN);  // enable the ADC

  // wait for voltages to become stable.
  delay(20);
}

void loop() {
}

// callback for received data
void receiveData(int byteCount){
  while(Wire.available()) {
    int command = Wire.read();
    int cmdmask = command & CMD_MASK;
    if(command==CMD_CPU_TEMP) {
        cputemp();
    } else if(command==CMD_WIND_DIR) {
      // Wind direction
      intout(0);
    } else if (cmdmask==CMD_WIND) {
      windout(command);
    } else if(cmdmask==CMD_WIND_GUST) {
      gustout(command);
    } else if(cmdmask==CMD_RAIN) {
      rainout(command);
    }
  }
}

// callback for sending data
void sendData(){
  Wire.write(output,output_size);
}

// utility to send a single byte
void byteout(byte val) {
  output_size = 1;
  output[0] = val;
}

// utility to send a 16 bit little endian integer
void intout(int val) {
  output_size = 2;
  output[0] = val & 0xff;
  output[1] = (val>>8) & 0xff;
}

// Utility to send a double value as a 16 bit little endian integer
void floatout(float val) {
  intout((int)(val*100));
}

// Utility to send a double value as a 16 bit little endian integer
void doubleout(double val) {
  intout((int)(val*100));
}

/*
 * cpu temperature sensor
 */
void cputemp() {
  // Start the ADC
  ADCSRA |= _BV(ADSC);

  // Detect end-of-conversion
  while (bit_is_set(ADCSRA,ADSC));

  // Reading register "ADCW" takes care of how to read ADCL and ADCH.
  unsigned int wADC = ADCW;

  // The offset of 324.31 could be wrong. It is just an indication.
  // We are using unit 0.01 on the pi so multiply by 100 here
  intout( (int)(100*(wADC - 324.31 ) / 1.22) );
}

