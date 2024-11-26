
"""Upload weather data to piweather.center.

This service enables pywws to upload data to our weather station.

* Install this file into modules (e.g. ~/weather/modules)
* Example ``weather.ini`` configuration::

    [piweather]
    url = http://127.0.0.1:8081/api/inbound/pywws
    pass key = secret

    [logged]
    services = ['piweather', 'metoffice', 'underground']

    [live]
    services = ['piweather', 'metoffice', 'underground']

* Add appropriate config for weatheringress:

    # station id
    home:
      sensors:
        # pywws "sensor" which will appear in the system
        # as home.pywws
        pywws:
          name: "Pywws Example station"
          source:
            # this service submits to weatheringress using http
            http:
              # http method used
              method: "GET"
              # path within weatheringress, so this will then become /api/inbound/pywws
              path: "pywws"
              # secret key to try to limit requests from just this station
              passKey: "secret"
          # metrics are passed in the query string
          format: "query"
          # the timestamp is from the dateutc parameter
          timestamp: "dateutc"
          # The readings from pywws
          readings:
            # outdoor temperature, so this becomes home.pywws.temp in this example
            temp:
              source: "tempout"
              type: "celsius"
            # indoor temperature
            tempin:
              source: "tempin"
              type: "celsius"
            # indoor humidity
            humidityin:
              source: "humin"
              type: "RelativeHumidity"
            # outdoor humidity
            humidity:
              source: "humout"
              type: "RelativeHumidity"
            # relative pressure
            pressure:
              source: "pressurerel"
              type: "PressureHPA"
            # wind direction
            winddir:
              source: "winddir"
              type: "degree"
            # average windspeed
            windspeed:
              source: "windspeed"
              type: "meterspersecond"
            # wind gust
            windgust:
              source: "windgust"
              type: "meterspersecond"
            # rain fall
            rain:
              source: "rain"
              type: "millimeters"
            # rain fall today
            rainday:
              source: "rainday"
              type: "millimeters"
            # rain fall in last 24 hours
            rain24:
              source: "rain24"
              type: "millimeters"
            # dewpoint
            dewpoint:
              source: "dewpoint"
              type: "celsius"
            # windchill
            windchill:
              source: "windchill"
              type: "celsius"

"""

from __future__ import absolute_import, unicode_literals

from contextlib import contextmanager
from datetime import timedelta
import logging
import os
import sys

import requests

import pywws
from pywws.conversions import rain_inch
from pywws.process import get_day_end_hour
import pywws.service
from pywws.timezone import time_zone

__docformat__ = "restructuredtext en"
service_name = os.path.splitext(os.path.basename(__file__))[0]
logger = logging.getLogger(__name__)


class ToService(pywws.service.CatchupDataService):
    config = {
        'url'     : ('', False, None),
        'pass key': ('', True,  'passKey'),
        }
    fixed_data = {'softwaretype': 'pywws-' + pywws.__version__}
    interval = timedelta(seconds=40)
    logger = logger
    service_name = service_name
    template = """
#live#
#idx          "'dateutc'     : '%Y-%m-%d %H:%M:%S',"#
#wind_dir     '"winddir": "%.d",' '' 'winddir_degrees(x)'#
#wind_ave     '"windspeed"    : "%.2f",'#
#wind_gust    '"windgust"   : "%.2f",'#
#calc 'wind_chill(data["temp_out"],data["wind_ave"])'         '"windchill" : "%.1f",'#
#calc 'dew_point(data["temp_out"],data["hum_out"])'           '"dewpoint" : "%.1f",'#
#hum_out      '"humout"     : "%.d",'#
#hum_in       '"humin"      : "%.d",'#
#temp_in      '"tempin"   : "%.1f",'#
#temp_out     '"tempout"  : "%.1f",'#
#rel_pressure '"pressurerel" : "%.1f",'#
#abs_pressure '"pressureabs" : "%.1f",'#
#rain         '"rain"     : "%.1f",'#
#calc 'rain_hour(data)'            '"rainhr": "%.1f",'#
#calc 'rain_24hr(data)'            '"rain24": "%.1f",'#
#calc 'rain_day(data)'             '"rainday": "%.1f",'#
"""

    def __init__(self, context, check_params=True):
        super(ToService, self).__init__(context, check_params)
        logger.info('url ' +self.params['url'])

    @contextmanager
    def session(self):
        with requests.Session() as session:
            yield session, 'OK'

    def valid_data(self, data):
        return any([data[x] is not None for x in (
            'wind_dir', 'wind_ave', 'wind_gust', 'hum_out', 'temp_out',
            'rel_pressure')])

    def upload_data(self, session, prepared_data={}):
        try:
            rsp = session.get(self.params['url'],
                              params=prepared_data, timeout=60)
        except Exception as ex:
            return False, repr(ex)
        if rsp.status_code != 200:
            return False, 'http status: {:d}'.format(rsp.status_code)
        return True, 'OK'


if __name__ == "__main__":
    sys.exit(pywws.service.main(ToService))
