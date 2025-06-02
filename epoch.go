package main

/**
 * Number of milliseconds in one day.
 */
const MILLIS_IN_A_DAY = 86400000.0

	 /**
	  * Converts a TLE epoch year and fractional Julian day into a millisecond epoch from January 1,
	  * 1970 00:00:00.
	  *
	  * @param year the year to be converted
	  * @param julianDay the fractional Julian day to be converted
	  * @return the millisecond epoch
	  */
	 static long toMillisecondEpoch(int year, double julianDay) {
		 Calendar calendar = Calendar.getInstance();
		 calendar.clear();
		 calendar.setTimeZone(UTC_TIME_ZONE);
 
		 // Set year
		 calendar.set(Calendar.YEAR, year);
 
		 // Set day
		 int wholeDay = (int) julianDay;
		 calendar.set(Calendar.DAY_OF_YEAR, wholeDay);
 
		 // Set millisecond
		 double dayFraction = julianDay - (double) wholeDay;
		 BigDecimal ms = new BigDecimal(dayFraction * MILLIS_IN_A_DAY);
		 int millisecond = ms.setScale(0, RoundingMode.HALF_UP).intValueExact();
		 calendar.set(Calendar.MILLISECOND, millisecond);
 
		 return calendar.getTimeInMillis();
	 }