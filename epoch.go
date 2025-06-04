package main

import "log"

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
	func toMillisecondEpoch(year int, julianDay float64) int64 {
		return 0
	}

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

	  /**
     * Formats a TLE with the millisecond from January 1, 1970 00:00:00.
     *
     * @param epochMillisecond the epoch millisecond to be formatted
     * @return the formatted epoch
     */
	 static String formatForTLE(long epochMillisecond) {
        // Get the year and day
        Calendar calendar = Calendar.getInstance();
        calendar.clear();
        calendar.setTimeZone(UTC_TIME_ZONE);
        calendar.setTimeInMillis(epochMillisecond);
        int year = calendar.get(Calendar.YEAR);
        int dayOfYear = calendar.get(Calendar.DAY_OF_YEAR);
        int twoDigitYear = year % 100;

        // Reset the calendar to contain only milliseconds accounted for by year and day of year
        calendar.clear();
        calendar.setTimeZone(UTC_TIME_ZONE);
        calendar.set(Calendar.YEAR, year);
        calendar.set(Calendar.DAY_OF_YEAR, dayOfYear);

        // Get the fractional part of the day by subtracting out the year and day of year
        long remainingMilliseconds = epochMillisecond - calendar.getTimeInMillis();
        double fractionalDay = remainingMilliseconds / MILLIS_IN_A_DAY;

        // Format year and day for TLE epoch field
        String decimal = DECIMAL_FORMAT_ATOMIC_REFERENCE.get().format(fractionalDay);
        return String.format("%02d%3d%-8s", twoDigitYear, dayOfYear, decimal);
    }

    /**
     * Extracts the epoch year from the epoch millisecond.
     *
     * @param epochMillisecond the number of milliseconds since January 1, 1970 00:00:00
     * @return the year
     */
	 static int getEpochYear(long epochMillisecond) {
        Calendar calendar = Calendar.getInstance();
        calendar.clear();
        calendar.setTimeZone(UTC_TIME_ZONE);
        calendar.setTimeInMillis(epochMillisecond);
        return calendar.get(Calendar.YEAR);
    }


   /**
     * Extracts the fractional Julian day from the epoch millisecond.
     *
     * @param epochMillisecond the number of milliseconds since January 1, 1970 00:00:00
     * @return the fractional Julian day
     */
    static double getEpochJulianDay(long epochMillisecond) {
        // Get the year and day
        Calendar calendar = Calendar.getInstance();
        calendar.clear();
        calendar.setTimeZone(UTC_TIME_ZONE);
        calendar.setTimeInMillis(epochMillisecond);
        int year = calendar.get(Calendar.YEAR);
        int dayOfYear = calendar.get(Calendar.DAY_OF_YEAR);

        // Get a calendar that contains milliseconds accounted for by the year and day of year
        calendar.clear();
        calendar.setTimeZone(UTC_TIME_ZONE);
        calendar.set(Calendar.YEAR, year);
        calendar.set(Calendar.DAY_OF_YEAR, dayOfYear);

        // Get the fractional part of the day by subtracting out the year and day of year
        long remainingMilliseconds = epochMillisecond - calendar.getTimeInMillis();
        double fractionalDay = remainingMilliseconds / MILLIS_IN_A_DAY;

        // Format year and day for TLE epoch field
        return (double) dayOfYear + fractionalDay;
    }