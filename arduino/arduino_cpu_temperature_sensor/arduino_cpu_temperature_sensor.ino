/*
 An example Arduino i2c slave for the pi weather station.
 
 In this example we create an i2c slave at address 0x04.
 
 It has one single command, 0x02 which will read the arduino's CPU Temperature.
 
 On the Raspbery PI, we need to create the sensor, so create /etc/weather/arduino
 with the following:
 
[arduino/cpu]
title                   Arduino CPU Temperature
frequency               30
format                  Ard %.2fC
unit                    0.01
sensor-type             i2c
i2c-address             04
i2c-command             02
i2c-response-type       word
i2c-rw-delay            10000
i2c-post-delay          10000

  Here we defing a sensor which will run once every 30 seconds.
  It uses i2c and expects a word back. The unit 0.01 means we are returning the
  temperature in 100'ths of a C (the *100 in the code below).
  
  Note: i2c-rw-delay tells our i2c library to wait 10ms between the write and
  the read from the arduino. This is to allow the arduino to perform the actual
  reading as part of this request, If you already have the info you could reduce
  this.
  
  The i2c-post-delay is also 10ms and is a delay after reading before allowing
  the next command to be sent. This is because I have managed in original testing
  to crash the arduino by sending commands too fast, so this just keeps multiple
  sensors on this slave to be delayed slightly, preventing this.
  
 */
#include <Wire.h>

#define SLAVE_ADDRESS 0x04

byte output[2];

void setup() {
  // initialize i2c as slave
  Wire.begin(SLAVE_ADDRESS);
  
  // define callbacks for i2c communication
  Wire.onReceive(receiveData);
  Wire.onRequest(sendData);
  
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
    
    if(command==2) {
      // Start the ADC
      ADCSRA |= _BV(ADSC);

      // Detect end-of-conversion
      while (bit_is_set(ADCSRA,ADSC));

      // Reading register "ADCW" takes care of how to read ADCL and ADCH.
      unsigned int wADC = ADCW;

      // The offset of 324.31 could be wrong. It is just an indication.
      // We are using unit 0.01 on the pi so multiply by 100 here
      int temp = (int)(100*(wADC - 324.31 ) / 1.22);
      
      // We are using response type of word, so it's 2 bytes, lsb in the response
      output[0] = temp & 0xff;
      output[1] = (temp>>8) & 0xff;
      
    } else {
      // Dummy output, respond with 0
      output[0]=0;
      output[1]=0;
    }
  }
}

// callback for sending data
void sendData(){
  Wire.write(output,2);
}

