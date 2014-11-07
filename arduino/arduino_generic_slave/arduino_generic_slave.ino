/*
 A generic Arduinmo i2c slave for the pi weather station project.
 
 This sketch enables an Arduino UNO to behave as an i2c slave to
 the Raspberry PI. It provides generic access to the arduino's GPIO ports
 from code running on the PI including ADC's and configuring the various
 I/O lines.
 
 Note: On the UNO ADC 4 & 5 are not available as they are used for i2C communication.
       If you try to use them we will return 0. Same for any ADC channels not available
       on the platform.

 If you have multiple arduino's on the bus then you must select a unique address for each one. 
 
 Supported commands: Note: Any command not listed here will return a single byte value 0
 
 Command  Description                                    Returns          Unit        Since
 ==========================================================================================
 0x00     Returns firmware version number                word             0.01        1
 0x02     Arduino CPU Temperature                        word             0.01        1
 0x03     Turn off onboard LED                           byte  always 0   1           1
 0x04     Turn on onboard LED                            byte  always 0   1           1
 0x1x     Read analogue channel x                        word             1*          1

 * for analogue read the raw value of the channel is returned. unit & other conversions should
   be one on the client side 
  
 */
#include <Wire.h>

// The slave address
#define SLAVE_ADDRESS 0x04

// Firmware version, returned by command 0x00. This should be incremented if we change the
// supported command set.
// Version 0 was the initial demonstration slave
// Version 1 is the first version
#define VERSION 1

int output_size;
byte output[2];

void setup() {
  // initialize i2c as slave
  Wire.begin(SLAVE_ADDRESS);
  
  // define callbacks for i2c communication
  Wire.onReceive(receiveData);
  Wire.onRequest(sendData);

  // Set onboard LED
  pinMode(LED_BUILTIN,OUTPUT);
  
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
    int major = command & 0xf0;
    int minor = command &0x0f;
    
    switch(major) {
      case 0x00:
        control(minor);
        break;
      case 0x10:
        adc(minor);
        break;
      
      default:
        // Dummy output, respond with 0
        byteout(0);
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

// Handles commands 0x00-0x0f
void control(int minor) {
  switch(minor) {
    // Firmware version
    case 0x00:
      intout(VERSION);
      break;
    // CPU Temperature
    case 0x02:
      cputemp();
      break;
    // Builtin LED off
    case 0x03:
      digitalWrite(LED_BUILTIN,LOW);
      byteout(0);
      break;
    // Builtin LED on
    case 0x04:
      digitalWrite(LED_BUILTIN,HIGH);
      byteout(0);
      break;
    // default do nothing
    default:
      byteout(0);
      break;
  }
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

/*
 * adc - if an invalid port for this platform then returns 0
 * i.e. uno shares i2c with analogue so adc 4 & 5 are also invalid
 */
void adc(int channel) {
  int pin=analogInputToDigitalPin(channel);
  if(pin==-1 || pin==SDA || pin==SCL)
    intout(0);
  else
    intout(analogRead(channel));
}

