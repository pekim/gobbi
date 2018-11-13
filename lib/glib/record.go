// This is a generated file - DO NOT EDIT

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Contains the public fields of a GArray.
/*

C type

GArray
*/
type Array struct {
	native *C.GArray
	Data   string
	Len    uint32
}

func ArrayNewFromC(u unsafe.Pointer) *Array {
	c := (*C.GArray)(u)
	if c == nil {
		return nil
	}

	g := &Array{
		Data:   C.GoString(c.data),
		Len:    (uint32)(c.len),
		native: c,
	}

	return g
}

func (recv *Array) ToC() unsafe.Pointer {
	recv.native.data =
		C.CString(recv.Data)
	recv.native.len =
		(C.guint)(recv.Len)

	return (unsafe.Pointer)(recv.native)
}

// The GAsyncQueue struct is an opaque data structure which represents
// an asynchronous queue. It should only be accessed through the
// g_async_queue_* functions.
/*

C type

GAsyncQueue
*/
type AsyncQueue struct {
	native *C.GAsyncQueue
}

func AsyncQueueNewFromC(u unsafe.Pointer) *AsyncQueue {
	c := (*C.GAsyncQueue)(u)
	if c == nil {
		return nil
	}

	g := &AsyncQueue{native: c}

	return g
}

func (recv *AsyncQueue) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Returns the length of the queue.
//
// Actually this function returns the number of data items in
// the queue minus the number of waiting threads, so a negative
// value means waiting threads, and a positive value means available
// entries in the @queue. A return value of 0 could mean n entries
// in the queue and n threads waiting. This can happen due to locking
// of the queue or due to scheduling.
/*

C function

g_async_queue_length
*/
func (recv *AsyncQueue) Length() int32 {
	retC := C.g_async_queue_length((*C.GAsyncQueue)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the length of the queue.
//
// Actually this function returns the number of data items in
// the queue minus the number of waiting threads, so a negative
// value means waiting threads, and a positive value means available
// entries in the @queue. A return value of 0 could mean n entries
// in the queue and n threads waiting. This can happen due to locking
// of the queue or due to scheduling.
//
// This function must be called while holding the @queue's lock.
/*

C function

g_async_queue_length_unlocked
*/
func (recv *AsyncQueue) LengthUnlocked() int32 {
	retC := C.g_async_queue_length_unlocked((*C.GAsyncQueue)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Acquires the @queue's lock. If another thread is already
// holding the lock, this call will block until the lock
// becomes available.
//
// Call g_async_queue_unlock() to drop the lock again.
//
// While holding the lock, you can only call the
// g_async_queue_*_unlocked() functions on @queue. Otherwise,
// deadlock may occur.
/*

C function

g_async_queue_lock
*/
func (recv *AsyncQueue) Lock() {
	C.g_async_queue_lock((*C.GAsyncQueue)(recv.native))

	return
}

// Pops data from the @queue. If @queue is empty, this function
// blocks until data becomes available.
/*

C function

g_async_queue_pop
*/
func (recv *AsyncQueue) Pop() uintptr {
	retC := C.g_async_queue_pop((*C.GAsyncQueue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Pops data from the @queue. If @queue is empty, this function
// blocks until data becomes available.
//
// This function must be called while holding the @queue's lock.
/*

C function

g_async_queue_pop_unlocked
*/
func (recv *AsyncQueue) PopUnlocked() uintptr {
	retC := C.g_async_queue_pop_unlocked((*C.GAsyncQueue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Pushes the @data into the @queue. @data must not be %NULL.
/*

C function

g_async_queue_push
*/
func (recv *AsyncQueue) Push(data uintptr) {
	c_data := (C.gpointer)(data)

	C.g_async_queue_push((*C.GAsyncQueue)(recv.native), c_data)

	return
}

// Pushes the @data into the @queue. @data must not be %NULL.
//
// This function must be called while holding the @queue's lock.
/*

C function

g_async_queue_push_unlocked
*/
func (recv *AsyncQueue) PushUnlocked(data uintptr) {
	c_data := (C.gpointer)(data)

	C.g_async_queue_push_unlocked((*C.GAsyncQueue)(recv.native), c_data)

	return
}

// Increases the reference count of the asynchronous @queue by 1.
// You do not need to hold the lock to call this function.
/*

C function

g_async_queue_ref
*/
func (recv *AsyncQueue) Ref() *AsyncQueue {
	retC := C.g_async_queue_ref((*C.GAsyncQueue)(recv.native))
	retGo := AsyncQueueNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Increases the reference count of the asynchronous @queue by 1.
/*

C function

g_async_queue_ref_unlocked
*/
func (recv *AsyncQueue) RefUnlocked() {
	C.g_async_queue_ref_unlocked((*C.GAsyncQueue)(recv.native))

	return
}

// Pops data from the @queue. If the queue is empty, blocks until
// @end_time or until data becomes available.
//
// If no data is received before @end_time, %NULL is returned.
//
// To easily calculate @end_time, a combination of g_get_current_time()
// and g_time_val_add() can be used.
/*

C function

g_async_queue_timed_pop
*/
func (recv *AsyncQueue) TimedPop(endTime *TimeVal) uintptr {
	c_end_time := (*C.GTimeVal)(C.NULL)
	if endTime != nil {
		c_end_time = (*C.GTimeVal)(endTime.ToC())
	}

	retC := C.g_async_queue_timed_pop((*C.GAsyncQueue)(recv.native), c_end_time)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Pops data from the @queue. If the queue is empty, blocks until
// @end_time or until data becomes available.
//
// If no data is received before @end_time, %NULL is returned.
//
// To easily calculate @end_time, a combination of g_get_current_time()
// and g_time_val_add() can be used.
//
// This function must be called while holding the @queue's lock.
/*

C function

g_async_queue_timed_pop_unlocked
*/
func (recv *AsyncQueue) TimedPopUnlocked(endTime *TimeVal) uintptr {
	c_end_time := (*C.GTimeVal)(C.NULL)
	if endTime != nil {
		c_end_time = (*C.GTimeVal)(endTime.ToC())
	}

	retC := C.g_async_queue_timed_pop_unlocked((*C.GAsyncQueue)(recv.native), c_end_time)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Pops data from the @queue. If the queue is empty, blocks for
// @timeout microseconds, or until data becomes available.
//
// If no data is received before the timeout, %NULL is returned.
/*

C function

g_async_queue_timeout_pop
*/
func (recv *AsyncQueue) TimeoutPop(timeout uint64) uintptr {
	c_timeout := (C.guint64)(timeout)

	retC := C.g_async_queue_timeout_pop((*C.GAsyncQueue)(recv.native), c_timeout)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Pops data from the @queue. If the queue is empty, blocks for
// @timeout microseconds, or until data becomes available.
//
// If no data is received before the timeout, %NULL is returned.
//
// This function must be called while holding the @queue's lock.
/*

C function

g_async_queue_timeout_pop_unlocked
*/
func (recv *AsyncQueue) TimeoutPopUnlocked(timeout uint64) uintptr {
	c_timeout := (C.guint64)(timeout)

	retC := C.g_async_queue_timeout_pop_unlocked((*C.GAsyncQueue)(recv.native), c_timeout)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Tries to pop data from the @queue. If no data is available,
// %NULL is returned.
/*

C function

g_async_queue_try_pop
*/
func (recv *AsyncQueue) TryPop() uintptr {
	retC := C.g_async_queue_try_pop((*C.GAsyncQueue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Tries to pop data from the @queue. If no data is available,
// %NULL is returned.
//
// This function must be called while holding the @queue's lock.
/*

C function

g_async_queue_try_pop_unlocked
*/
func (recv *AsyncQueue) TryPopUnlocked() uintptr {
	retC := C.g_async_queue_try_pop_unlocked((*C.GAsyncQueue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Releases the queue's lock.
//
// Calling this function when you have not acquired
// the with g_async_queue_lock() leads to undefined
// behaviour.
/*

C function

g_async_queue_unlock
*/
func (recv *AsyncQueue) Unlock() {
	C.g_async_queue_unlock((*C.GAsyncQueue)(recv.native))

	return
}

// Decreases the reference count of the asynchronous @queue by 1.
//
// If the reference count went to 0, the @queue will be destroyed
// and the memory allocated will be freed. So you are not allowed
// to use the @queue afterwards, as it might have disappeared.
// You do not need to hold the lock to call this function.
/*

C function

g_async_queue_unref
*/
func (recv *AsyncQueue) Unref() {
	C.g_async_queue_unref((*C.GAsyncQueue)(recv.native))

	return
}

// Decreases the reference count of the asynchronous @queue by 1
// and releases the lock. This function must be called while holding
// the @queue's lock. If the reference count went to 0, the @queue
// will be destroyed and the memory allocated will be freed.
/*

C function

g_async_queue_unref_and_unlock
*/
func (recv *AsyncQueue) UnrefAndUnlock() {
	C.g_async_queue_unref_and_unlock((*C.GAsyncQueue)(recv.native))

	return
}

// The `GBookmarkFile` structure contains only
// private data and should not be directly accessed.
/*

C type

GBookmarkFile
*/
type BookmarkFile struct {
	native *C.GBookmarkFile
}

func BookmarkFileNewFromC(u unsafe.Pointer) *BookmarkFile {
	c := (*C.GBookmarkFile)(u)
	if c == nil {
		return nil
	}

	g := &BookmarkFile{native: c}

	return g
}

func (recv *BookmarkFile) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : GByteArray

// The #GCond struct is an opaque data structure that represents a
// condition. Threads can block on a #GCond if they find a certain
// condition to be false. If other threads change the state of this
// condition they signal the #GCond, and that causes the waiting
// threads to be woken up.
//
// Consider the following example of a shared variable.  One or more
// threads can wait for data to be published to the variable and when
// another thread publishes the data, it can signal one of the waiting
// threads to wake up to collect the data.
//
// Here is an example for using GCond to block a thread until a condition
// is satisfied:
// |[<!-- language="C" -->
// gpointer current_data = NULL;
// GMutex data_mutex;
// GCond data_cond;
//
// void
// push_data (gpointer data)
// {
// g_mutex_lock (&data_mutex);
// current_data = data;
// g_cond_signal (&data_cond);
// g_mutex_unlock (&data_mutex);
// }
//
// gpointer
// pop_data (void)
// {
// gpointer data;
//
// g_mutex_lock (&data_mutex);
// while (!current_data)
// g_cond_wait (&data_cond, &data_mutex);
// data = current_data;
// current_data = NULL;
// g_mutex_unlock (&data_mutex);
//
// return data;
// }
// ]|
// Whenever a thread calls pop_data() now, it will wait until
// current_data is non-%NULL, i.e. until some other thread
// has called push_data().
//
// The example shows that use of a condition variable must always be
// paired with a mutex.  Without the use of a mutex, there would be a
// race between the check of @current_data by the while loop in
// pop_data() and waiting. Specifically, another thread could set
// @current_data after the check, and signal the cond (with nobody
// waiting on it) before the first thread goes to sleep. #GCond is
// specifically useful for its ability to release the mutex and go
// to sleep atomically.
//
// It is also important to use the g_cond_wait() and g_cond_wait_until()
// functions only inside a loop which checks for the condition to be
// true.  See g_cond_wait() for an explanation of why the condition may
// not be true even after it returns.
//
// If a #GCond is allocated in static storage then it can be used
// without initialisation.  Otherwise, you should call g_cond_init()
// on it and g_cond_clear() when done.
//
// A #GCond should only be accessed via the g_cond_ functions.
/*

C type

GCond
*/
type Cond struct {
	native *C.GCond
	// Private : p
	// Private : i
}

func CondNewFromC(u unsafe.Pointer) *Cond {
	c := (*C.GCond)(u)
	if c == nil {
		return nil
	}

	g := &Cond{native: c}

	return g
}

func (recv *Cond) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// If threads are waiting for @cond, all of them are unblocked.
// If no threads are waiting for @cond, this function has no effect.
// It is good practice to lock the same mutex as the waiting threads
// while calling this function, though not required.
/*

C function

g_cond_broadcast
*/
func (recv *Cond) Broadcast() {
	C.g_cond_broadcast((*C.GCond)(recv.native))

	return
}

// If threads are waiting for @cond, at least one of them is unblocked.
// If no threads are waiting for @cond, this function has no effect.
// It is good practice to hold the same lock as the waiting thread
// while calling this function, though not required.
/*

C function

g_cond_signal
*/
func (recv *Cond) Signal() {
	C.g_cond_signal((*C.GCond)(recv.native))

	return
}

// Unsupported : g_cond_wait : unsupported parameter mutex : no type generator for Mutex (GMutex*) for param mutex

// The #GData struct is an opaque data structure to represent a
// [Keyed Data List][glib-Keyed-Data-Lists]. It should only be
// accessed via the following functions.
/*

C type

GData
*/
type Data struct {
	native *C.GData
}

func DataNewFromC(u unsafe.Pointer) *Data {
	c := (*C.GData)(u)
	if c == nil {
		return nil
	}

	g := &Data{native: c}

	return g
}

func (recv *Data) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Represents a day between January 1, Year 1 and a few thousand years in
// the future. None of its members should be accessed directly.
//
// If the #GDate-struct is obtained from g_date_new(), it will be safe
// to mutate but invalid and thus not safe for calendrical computations.
//
// If it's declared on the stack, it will contain garbage so must be
// initialized with g_date_clear(). g_date_clear() makes the date invalid
// but sane. An invalid date doesn't represent a day, it's "empty." A date
// becomes valid after you set it to a Julian day or you set a day, month,
// and year.
/*

C type

GDate
*/
type Date struct {
	native *C.GDate
	// Bitfield not supported : 32 julian_days
	// Bitfield not supported :  1 julian
	// Bitfield not supported :  1 dmy
	// Bitfield not supported :  6 day
	// Bitfield not supported :  4 month
	// Bitfield not supported : 16 year
}

func DateNewFromC(u unsafe.Pointer) *Date {
	c := (*C.GDate)(u)
	if c == nil {
		return nil
	}

	g := &Date{native: c}

	return g
}

func (recv *Date) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Allocates a #GDate and initializes
// it to a sane state. The new date will
// be cleared (as if you'd called g_date_clear()) but invalid (it won't
// represent an existing day). Free the return value with g_date_free().
/*

C function

g_date_new
*/
func DateNew() *Date {
	retC := C.g_date_new()
	retGo := DateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Like g_date_new(), but also sets the value of the date. Assuming the
// day-month-year triplet you pass in represents an existing day, the
// returned date will be valid.
/*

C function

g_date_new_dmy
*/
func DateNewDmy(day DateDay, month DateMonth, year DateYear) *Date {
	c_day := (C.GDateDay)(day)

	c_month := (C.GDateMonth)(month)

	c_year := (C.GDateYear)(year)

	retC := C.g_date_new_dmy(c_day, c_month, c_year)
	retGo := DateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Like g_date_new(), but also sets the value of the date. Assuming the
// Julian day number you pass in is valid (greater than 0, less than an
// unreasonably large number), the returned date will be valid.
/*

C function

g_date_new_julian
*/
func DateNewJulian(julianDay uint32) *Date {
	c_julian_day := (C.guint32)(julianDay)

	retC := C.g_date_new_julian(c_julian_day)
	retGo := DateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Increments a date some number of days.
// To move forward by weeks, add weeks*7 days.
// The date must be valid.
/*

C function

g_date_add_days
*/
func (recv *Date) AddDays(nDays uint32) {
	c_n_days := (C.guint)(nDays)

	C.g_date_add_days((*C.GDate)(recv.native), c_n_days)

	return
}

// Increments a date by some number of months.
// If the day of the month is greater than 28,
// this routine may change the day of the month
// (because the destination month may not have
// the current day in it). The date must be valid.
/*

C function

g_date_add_months
*/
func (recv *Date) AddMonths(nMonths uint32) {
	c_n_months := (C.guint)(nMonths)

	C.g_date_add_months((*C.GDate)(recv.native), c_n_months)

	return
}

// Increments a date by some number of years.
// If the date is February 29, and the destination
// year is not a leap year, the date will be changed
// to February 28. The date must be valid.
/*

C function

g_date_add_years
*/
func (recv *Date) AddYears(nYears uint32) {
	c_n_years := (C.guint)(nYears)

	C.g_date_add_years((*C.GDate)(recv.native), c_n_years)

	return
}

// If @date is prior to @min_date, sets @date equal to @min_date.
// If @date falls after @max_date, sets @date equal to @max_date.
// Otherwise, @date is unchanged.
// Either of @min_date and @max_date may be %NULL.
// All non-%NULL dates must be valid.
/*

C function

g_date_clamp
*/
func (recv *Date) Clamp(minDate *Date, maxDate *Date) {
	c_min_date := (*C.GDate)(C.NULL)
	if minDate != nil {
		c_min_date = (*C.GDate)(minDate.ToC())
	}

	c_max_date := (*C.GDate)(C.NULL)
	if maxDate != nil {
		c_max_date = (*C.GDate)(maxDate.ToC())
	}

	C.g_date_clamp((*C.GDate)(recv.native), c_min_date, c_max_date)

	return
}

// Initializes one or more #GDate structs to a sane but invalid
// state. The cleared dates will not represent an existing date, but will
// not contain garbage. Useful to init a date declared on the stack.
// Validity can be tested with g_date_valid().
/*

C function

g_date_clear
*/
func (recv *Date) Clear(nDates uint32) {
	c_n_dates := (C.guint)(nDates)

	C.g_date_clear((*C.GDate)(recv.native), c_n_dates)

	return
}

// qsort()-style comparison function for dates.
// Both dates must be valid.
/*

C function

g_date_compare
*/
func (recv *Date) Compare(rhs *Date) int32 {
	c_rhs := (*C.GDate)(C.NULL)
	if rhs != nil {
		c_rhs = (*C.GDate)(rhs.ToC())
	}

	retC := C.g_date_compare((*C.GDate)(recv.native), c_rhs)
	retGo := (int32)(retC)

	return retGo
}

// Computes the number of days between two dates.
// If @date2 is prior to @date1, the returned value is negative.
// Both dates must be valid.
/*

C function

g_date_days_between
*/
func (recv *Date) DaysBetween(date2 *Date) int32 {
	c_date2 := (*C.GDate)(C.NULL)
	if date2 != nil {
		c_date2 = (*C.GDate)(date2.ToC())
	}

	retC := C.g_date_days_between((*C.GDate)(recv.native), c_date2)
	retGo := (int32)(retC)

	return retGo
}

// Frees a #GDate returned from g_date_new().
/*

C function

g_date_free
*/
func (recv *Date) Free() {
	C.g_date_free((*C.GDate)(recv.native))

	return
}

// Returns the day of the month. The date must be valid.
/*

C function

g_date_get_day
*/
func (recv *Date) GetDay() DateDay {
	retC := C.g_date_get_day((*C.GDate)(recv.native))
	retGo := (DateDay)(retC)

	return retGo
}

// Returns the day of the year, where Jan 1 is the first day of the
// year. The date must be valid.
/*

C function

g_date_get_day_of_year
*/
func (recv *Date) GetDayOfYear() uint32 {
	retC := C.g_date_get_day_of_year((*C.GDate)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Returns the Julian day or "serial number" of the #GDate. The
// Julian day is simply the number of days since January 1, Year 1; i.e.,
// January 1, Year 1 is Julian day 1; January 2, Year 1 is Julian day 2,
// etc. The date must be valid.
/*

C function

g_date_get_julian
*/
func (recv *Date) GetJulian() uint32 {
	retC := C.g_date_get_julian((*C.GDate)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Returns the week of the year, where weeks are understood to start on
// Monday. If the date is before the first Monday of the year, return 0.
// The date must be valid.
/*

C function

g_date_get_monday_week_of_year
*/
func (recv *Date) GetMondayWeekOfYear() uint32 {
	retC := C.g_date_get_monday_week_of_year((*C.GDate)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Returns the month of the year. The date must be valid.
/*

C function

g_date_get_month
*/
func (recv *Date) GetMonth() DateMonth {
	retC := C.g_date_get_month((*C.GDate)(recv.native))
	retGo := (DateMonth)(retC)

	return retGo
}

// Returns the week of the year during which this date falls, if
// weeks are understood to begin on Sunday. The date must be valid.
// Can return 0 if the day is before the first Sunday of the year.
/*

C function

g_date_get_sunday_week_of_year
*/
func (recv *Date) GetSundayWeekOfYear() uint32 {
	retC := C.g_date_get_sunday_week_of_year((*C.GDate)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Returns the day of the week for a #GDate. The date must be valid.
/*

C function

g_date_get_weekday
*/
func (recv *Date) GetWeekday() DateWeekday {
	retC := C.g_date_get_weekday((*C.GDate)(recv.native))
	retGo := (DateWeekday)(retC)

	return retGo
}

// Returns the year of a #GDate. The date must be valid.
/*

C function

g_date_get_year
*/
func (recv *Date) GetYear() DateYear {
	retC := C.g_date_get_year((*C.GDate)(recv.native))
	retGo := (DateYear)(retC)

	return retGo
}

// Returns %TRUE if the date is on the first of a month.
// The date must be valid.
/*

C function

g_date_is_first_of_month
*/
func (recv *Date) IsFirstOfMonth() bool {
	retC := C.g_date_is_first_of_month((*C.GDate)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns %TRUE if the date is the last day of the month.
// The date must be valid.
/*

C function

g_date_is_last_of_month
*/
func (recv *Date) IsLastOfMonth() bool {
	retC := C.g_date_is_last_of_month((*C.GDate)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Checks if @date1 is less than or equal to @date2,
// and swap the values if this is not the case.
/*

C function

g_date_order
*/
func (recv *Date) Order(date2 *Date) {
	c_date2 := (*C.GDate)(C.NULL)
	if date2 != nil {
		c_date2 = (*C.GDate)(date2.ToC())
	}

	C.g_date_order((*C.GDate)(recv.native), c_date2)

	return
}

// Sets the day of the month for a #GDate. If the resulting
// day-month-year triplet is invalid, the date will be invalid.
/*

C function

g_date_set_day
*/
func (recv *Date) SetDay(day DateDay) {
	c_day := (C.GDateDay)(day)

	C.g_date_set_day((*C.GDate)(recv.native), c_day)

	return
}

// Sets the value of a #GDate from a day, month, and year.
// The day-month-year triplet must be valid; if you aren't
// sure it is, call g_date_valid_dmy() to check before you
// set it.
/*

C function

g_date_set_dmy
*/
func (recv *Date) SetDmy(day DateDay, month DateMonth, y DateYear) {
	c_day := (C.GDateDay)(day)

	c_month := (C.GDateMonth)(month)

	c_y := (C.GDateYear)(y)

	C.g_date_set_dmy((*C.GDate)(recv.native), c_day, c_month, c_y)

	return
}

// Sets the value of a #GDate from a Julian day number.
/*

C function

g_date_set_julian
*/
func (recv *Date) SetJulian(julianDate uint32) {
	c_julian_date := (C.guint32)(julianDate)

	C.g_date_set_julian((*C.GDate)(recv.native), c_julian_date)

	return
}

// Sets the month of the year for a #GDate.  If the resulting
// day-month-year triplet is invalid, the date will be invalid.
/*

C function

g_date_set_month
*/
func (recv *Date) SetMonth(month DateMonth) {
	c_month := (C.GDateMonth)(month)

	C.g_date_set_month((*C.GDate)(recv.native), c_month)

	return
}

// Parses a user-inputted string @str, and try to figure out what date it
// represents, taking the [current locale][setlocale] into account. If the
// string is successfully parsed, the date will be valid after the call.
// Otherwise, it will be invalid. You should check using g_date_valid()
// to see whether the parsing succeeded.
//
// This function is not appropriate for file formats and the like; it
// isn't very precise, and its exact behavior varies with the locale.
// It's intended to be a heuristic routine that guesses what the user
// means by a given string (and it does work pretty well in that
// capacity).
/*

C function

g_date_set_parse
*/
func (recv *Date) SetParse(str string) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	C.g_date_set_parse((*C.GDate)(recv.native), c_str)

	return
}

// Sets the value of a date from a #GTime value.
// The time to date conversion is done using the user's current timezone.
/*

C function

g_date_set_time
*/
func (recv *Date) SetTime(time Time) {
	c_time_ := (C.GTime)(time)

	C.g_date_set_time((*C.GDate)(recv.native), c_time_)

	return
}

// Sets the year for a #GDate. If the resulting day-month-year
// triplet is invalid, the date will be invalid.
/*

C function

g_date_set_year
*/
func (recv *Date) SetYear(year DateYear) {
	c_year := (C.GDateYear)(year)

	C.g_date_set_year((*C.GDate)(recv.native), c_year)

	return
}

// Moves a date some number of days into the past.
// To move by weeks, just move by weeks*7 days.
// The date must be valid.
/*

C function

g_date_subtract_days
*/
func (recv *Date) SubtractDays(nDays uint32) {
	c_n_days := (C.guint)(nDays)

	C.g_date_subtract_days((*C.GDate)(recv.native), c_n_days)

	return
}

// Moves a date some number of months into the past.
// If the current day of the month doesn't exist in
// the destination month, the day of the month
// may change. The date must be valid.
/*

C function

g_date_subtract_months
*/
func (recv *Date) SubtractMonths(nMonths uint32) {
	c_n_months := (C.guint)(nMonths)

	C.g_date_subtract_months((*C.GDate)(recv.native), c_n_months)

	return
}

// Moves a date some number of years into the past.
// If the current day doesn't exist in the destination
// year (i.e. it's February 29 and you move to a non-leap-year)
// then the day is changed to February 29. The date
// must be valid.
/*

C function

g_date_subtract_years
*/
func (recv *Date) SubtractYears(nYears uint32) {
	c_n_years := (C.guint)(nYears)

	C.g_date_subtract_years((*C.GDate)(recv.native), c_n_years)

	return
}

// Unsupported : g_date_to_struct_tm : unsupported parameter tm : no type generator for gpointer (tm*) for param tm

// Returns %TRUE if the #GDate represents an existing day. The date must not
// contain garbage; it should have been initialized with g_date_clear()
// if it wasn't allocated by one of the g_date_new() variants.
/*

C function

g_date_valid
*/
func (recv *Date) Valid() bool {
	retC := C.g_date_valid((*C.GDate)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Associates a string with a bit flag.
// Used in g_parse_debug_string().
/*

C type

GDebugKey
*/
type DebugKey struct {
	native *C.GDebugKey
	Key    string
	Value  uint32
}

func DebugKeyNewFromC(u unsafe.Pointer) *DebugKey {
	c := (*C.GDebugKey)(u)
	if c == nil {
		return nil
	}

	g := &DebugKey{
		Key:    C.GoString(c.key),
		Value:  (uint32)(c.value),
		native: c,
	}

	return g
}

func (recv *DebugKey) ToC() unsafe.Pointer {
	recv.native.key =
		C.CString(recv.Key)
	recv.native.value =
		(C.guint)(recv.Value)

	return (unsafe.Pointer)(recv.native)
}

// An opaque structure representing an opened directory.
/*

C type

GDir
*/
type Dir struct {
	native *C.GDir
}

func DirNewFromC(u unsafe.Pointer) *Dir {
	c := (*C.GDir)(u)
	if c == nil {
		return nil
	}

	g := &Dir{native: c}

	return g
}

func (recv *Dir) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Closes the directory and deallocates all related resources.
/*

C function

g_dir_close
*/
func (recv *Dir) Close() {
	C.g_dir_close((*C.GDir)(recv.native))

	return
}

// Retrieves the name of another entry in the directory, or %NULL.
// The order of entries returned from this function is not defined,
// and may vary by file system or other operating-system dependent
// factors.
//
// %NULL may also be returned in case of errors. On Unix, you can
// check `errno` to find out if %NULL was returned because of an error.
//
// On Unix, the '.' and '..' entries are omitted, and the returned
// name is in the on-disk encoding.
//
// On Windows, as is true of all GLib functions which operate on
// filenames, the returned name is in UTF-8.
/*

C function

g_dir_read_name
*/
func (recv *Dir) ReadName() string {
	retC := C.g_dir_read_name((*C.GDir)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Resets the given directory. The next call to g_dir_read_name()
// will return the first entry again.
/*

C function

g_dir_rewind
*/
func (recv *Dir) Rewind() {
	C.g_dir_rewind((*C.GDir)(recv.native))

	return
}

// The `GError` structure contains information about
// an error that has occurred.
/*

C type

GError
*/
type Error struct {
	native  *C.GError
	Domain  Quark
	Code    int32
	Message string
}

func ErrorNewFromC(u unsafe.Pointer) *Error {
	c := (*C.GError)(u)
	if c == nil {
		return nil
	}

	g := &Error{
		Code:    (int32)(c.code),
		Domain:  (Quark)(c.domain),
		Message: C.GoString(c.message),
		native:  c,
	}

	return g
}

func (recv *Error) ToC() unsafe.Pointer {
	recv.native.domain =
		(C.GQuark)(recv.Domain)
	recv.native.code =
		(C.gint)(recv.Code)
	recv.native.message =
		C.CString(recv.Message)

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_error_new : unsupported parameter ... : varargs

// Creates a new #GError; unlike g_error_new(), @message is
// not a printf()-style format string. Use this function if
// @message contains text you don't have control over,
// that could include printf() escape sequences.
/*

C function

g_error_new_literal
*/
func ErrorNewLiteral(domain Quark, code int32, message string) *Error {
	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	retC := C.g_error_new_literal(c_domain, c_code, c_message)
	retGo := ErrorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Makes a copy of @error.
/*

C function

g_error_copy
*/
func (recv *Error) Copy() *Error {
	retC := C.g_error_copy((*C.GError)(recv.native))
	retGo := ErrorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Frees a #GError and associated resources.
/*

C function

g_error_free
*/
func (recv *Error) Free() {
	C.g_error_free((*C.GError)(recv.native))

	return
}

// Returns %TRUE if @error matches @domain and @code, %FALSE
// otherwise. In particular, when @error is %NULL, %FALSE will
// be returned.
//
// If @domain contains a `FAILED` (or otherwise generic) error code,
// you should generally not check for it explicitly, but should
// instead treat any not-explicitly-recognized error code as being
// equivalent to the `FAILED` code. This way, if the domain is
// extended in the future to provide a more specific error code for
// a certain case, your code will still work.
/*

C function

g_error_matches
*/
func (recv *Error) Matches(domain Quark, code int32) bool {
	c_domain := (C.GQuark)(domain)

	c_code := (C.gint)(code)

	retC := C.g_error_matches((*C.GError)(recv.native), c_domain, c_code)
	retGo := retC == C.TRUE

	return retGo
}

// The #GHashTable struct is an opaque data structure to represent a
// [Hash Table][glib-Hash-Tables]. It should only be accessed via the
// following functions.
/*

C type

GHashTable
*/
type HashTable struct {
	native *C.GHashTable
}

func HashTableNewFromC(u unsafe.Pointer) *HashTable {
	c := (*C.GHashTable)(u)
	if c == nil {
		return nil
	}

	g := &HashTable{native: c}

	return g
}

func (recv *HashTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// A GHashTableIter structure represents an iterator that can be used
// to iterate over the elements of a #GHashTable. GHashTableIter
// structures are typically allocated on the stack and then initialized
// with g_hash_table_iter_init().
/*

C type

GHashTableIter
*/
type HashTableIter struct {
	native *C.GHashTableIter
	// Private : dummy1
	// Private : dummy2
	// Private : dummy3
	// Private : dummy4
	// Private : dummy5
	// Private : dummy6
}

func HashTableIterNewFromC(u unsafe.Pointer) *HashTableIter {
	c := (*C.GHashTableIter)(u)
	if c == nil {
		return nil
	}

	g := &HashTableIter{native: c}

	return g
}

func (recv *HashTableIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// The #GHook struct represents a single hook function in a #GHookList.
/*

C type

GHook
*/
type Hook struct {
	native *C.GHook
	Data   uintptr
	// next : record
	// prev : record
	RefCount uint32
	HookId   uint64
	Flags    uint32
	Func     uintptr
	// destroy : no type generator for DestroyNotify, GDestroyNotify
}

func HookNewFromC(u unsafe.Pointer) *Hook {
	c := (*C.GHook)(u)
	if c == nil {
		return nil
	}

	g := &Hook{
		Data:     (uintptr)(c.data),
		Flags:    (uint32)(c.flags),
		Func:     (uintptr)(c._func),
		HookId:   (uint64)(c.hook_id),
		RefCount: (uint32)(c.ref_count),
		native:   c,
	}

	return g
}

func (recv *Hook) ToC() unsafe.Pointer {
	recv.native.data =
		(C.gpointer)(recv.Data)
	recv.native.ref_count =
		(C.guint)(recv.RefCount)
	recv.native.hook_id =
		(C.gulong)(recv.HookId)
	recv.native.flags =
		(C.guint)(recv.Flags)
	recv.native._func =
		(C.gpointer)(recv.Func)

	return (unsafe.Pointer)(recv.native)
}

// Compares the ids of two #GHook elements, returning a negative value
// if the second id is greater than the first.
/*

C function

g_hook_compare_ids
*/
func (recv *Hook) CompareIds(sibling *Hook) int32 {
	c_sibling := (*C.GHook)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GHook)(sibling.ToC())
	}

	retC := C.g_hook_compare_ids((*C.GHook)(recv.native), c_sibling)
	retGo := (int32)(retC)

	return retGo
}

// The #GHookList struct represents a list of hook functions.
/*

C type

GHookList
*/
type HookList struct {
	native *C.GHookList
	SeqId  uint64
	// Bitfield not supported : 16 hook_size
	// Bitfield not supported :  1 is_setup
	// hooks : record
	Dummy3 uintptr
	// finalize_hook : no type generator for HookFinalizeFunc, GHookFinalizeFunc
	// no type for dummy
}

func HookListNewFromC(u unsafe.Pointer) *HookList {
	c := (*C.GHookList)(u)
	if c == nil {
		return nil
	}

	g := &HookList{
		Dummy3: (uintptr)(c.dummy3),
		SeqId:  (uint64)(c.seq_id),
		native: c,
	}

	return g
}

func (recv *HookList) ToC() unsafe.Pointer {
	recv.native.seq_id =
		(C.gulong)(recv.SeqId)
	recv.native.dummy3 =
		(C.gpointer)(recv.Dummy3)

	return (unsafe.Pointer)(recv.native)
}

// Removes all the #GHook elements from a #GHookList.
/*

C function

g_hook_list_clear
*/
func (recv *HookList) Clear() {
	C.g_hook_list_clear((*C.GHookList)(recv.native))

	return
}

// Initializes a #GHookList.
// This must be called before the #GHookList is used.
/*

C function

g_hook_list_init
*/
func (recv *HookList) Init(hookSize uint32) {
	c_hook_size := (C.guint)(hookSize)

	C.g_hook_list_init((*C.GHookList)(recv.native), c_hook_size)

	return
}

// Calls all of the #GHook functions in a #GHookList.
/*

C function

g_hook_list_invoke
*/
func (recv *HookList) Invoke(mayRecurse bool) {
	c_may_recurse :=
		boolToGboolean(mayRecurse)

	C.g_hook_list_invoke((*C.GHookList)(recv.native), c_may_recurse)

	return
}

// Calls all of the #GHook functions in a #GHookList.
// Any function which returns %FALSE is removed from the #GHookList.
/*

C function

g_hook_list_invoke_check
*/
func (recv *HookList) InvokeCheck(mayRecurse bool) {
	c_may_recurse :=
		boolToGboolean(mayRecurse)

	C.g_hook_list_invoke_check((*C.GHookList)(recv.native), c_may_recurse)

	return
}

// Unsupported : g_hook_list_marshal : unsupported parameter marshaller : no type generator for HookMarshaller (GHookMarshaller) for param marshaller

// Unsupported : g_hook_list_marshal_check : unsupported parameter marshaller : no type generator for HookCheckMarshaller (GHookCheckMarshaller) for param marshaller

// Blacklisted : GIConv

// Blacklisted : GIOChannel

// A table of functions used to handle different types of #GIOChannel
// in a generic way.
/*

C type

GIOFuncs
*/
type IOFuncs struct {
	native *C.GIOFuncs
	// no type for io_read
	// no type for io_write
	// no type for io_seek
	// no type for io_close
	// no type for io_create_watch
	// no type for io_free
	// no type for io_set_flags
	// no type for io_get_flags
}

func IOFuncsNewFromC(u unsafe.Pointer) *IOFuncs {
	c := (*C.GIOFuncs)(u)
	if c == nil {
		return nil
	}

	g := &IOFuncs{native: c}

	return g
}

func (recv *IOFuncs) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// The GKeyFile struct contains only private data
// and should not be accessed directly.
/*

C type

GKeyFile
*/
type KeyFile struct {
	native *C.GKeyFile
}

func KeyFileNewFromC(u unsafe.Pointer) *KeyFile {
	c := (*C.GKeyFile)(u)
	if c == nil {
		return nil
	}

	g := &KeyFile{native: c}

	return g
}

func (recv *KeyFile) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// The #GList struct is used for each element in a doubly-linked list.
/*

C type

GList
*/
type List struct {
	native *C.GList
	Data   uintptr
	// next : record
	// prev : record
}

func ListNewFromC(u unsafe.Pointer) *List {
	c := (*C.GList)(u)
	if c == nil {
		return nil
	}

	g := &List{
		Data:   (uintptr)(c.data),
		native: c,
	}

	return g
}

func (recv *List) ToC() unsafe.Pointer {
	recv.native.data =
		(C.gpointer)(recv.Data)

	return (unsafe.Pointer)(recv.native)
}

// The `GMainContext` struct is an opaque data
// type representing a set of sources to be handled in a main loop.
/*

C type

GMainContext
*/
type MainContext struct {
	native *C.GMainContext
}

func MainContextNewFromC(u unsafe.Pointer) *MainContext {
	c := (*C.GMainContext)(u)
	if c == nil {
		return nil
	}

	g := &MainContext{native: c}

	return g
}

func (recv *MainContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Creates a new #GMainContext structure.
/*

C function

g_main_context_new
*/
func MainContextNew() *MainContext {
	retC := C.g_main_context_new()
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Tries to become the owner of the specified context.
// If some other thread is the owner of the context,
// returns %FALSE immediately. Ownership is properly
// recursive: the owner can require ownership again
// and will release ownership when g_main_context_release()
// is called as many times as g_main_context_acquire().
//
// You must be the owner of a context before you
// can call g_main_context_prepare(), g_main_context_query(),
// g_main_context_check(), g_main_context_dispatch().
/*

C function

g_main_context_acquire
*/
func (recv *MainContext) Acquire() bool {
	retC := C.g_main_context_acquire((*C.GMainContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Adds a file descriptor to the set of file descriptors polled for
// this context. This will very seldom be used directly. Instead
// a typical event source will use g_source_add_unix_fd() instead.
/*

C function

g_main_context_add_poll
*/
func (recv *MainContext) AddPoll(fd *PollFD, priority int32) {
	c_fd := (*C.GPollFD)(C.NULL)
	if fd != nil {
		c_fd = (*C.GPollFD)(fd.ToC())
	}

	c_priority := (C.gint)(priority)

	C.g_main_context_add_poll((*C.GMainContext)(recv.native), c_fd, c_priority)

	return
}

// Unsupported : g_main_context_check : unsupported parameter fds :

// Dispatches all pending sources.
//
// You must have successfully acquired the context with
// g_main_context_acquire() before you may call this function.
/*

C function

g_main_context_dispatch
*/
func (recv *MainContext) Dispatch() {
	C.g_main_context_dispatch((*C.GMainContext)(recv.native))

	return
}

// Finds a source with the given source functions and user data.  If
// multiple sources exist with the same source function and user data,
// the first one found will be returned.
/*

C function

g_main_context_find_source_by_funcs_user_data
*/
func (recv *MainContext) FindSourceByFuncsUserData(funcs *SourceFuncs, userData uintptr) *Source {
	c_funcs := (*C.GSourceFuncs)(C.NULL)
	if funcs != nil {
		c_funcs = (*C.GSourceFuncs)(funcs.ToC())
	}

	c_user_data := (C.gpointer)(userData)

	retC := C.g_main_context_find_source_by_funcs_user_data((*C.GMainContext)(recv.native), c_funcs, c_user_data)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Finds a #GSource given a pair of context and ID.
//
// It is a programmer error to attempt to lookup a non-existent source.
//
// More specifically: source IDs can be reissued after a source has been
// destroyed and therefore it is never valid to use this function with a
// source ID which may have already been removed.  An example is when
// scheduling an idle to run in another thread with g_idle_add(): the
// idle may already have run and been removed by the time this function
// is called on its (now invalid) source ID.  This source ID may have
// been reissued, leading to the operation being performed against the
// wrong source.
/*

C function

g_main_context_find_source_by_id
*/
func (recv *MainContext) FindSourceById(sourceId uint32) *Source {
	c_source_id := (C.guint)(sourceId)

	retC := C.g_main_context_find_source_by_id((*C.GMainContext)(recv.native), c_source_id)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Finds a source with the given user data for the callback.  If
// multiple sources exist with the same user data, the first
// one found will be returned.
/*

C function

g_main_context_find_source_by_user_data
*/
func (recv *MainContext) FindSourceByUserData(userData uintptr) *Source {
	c_user_data := (C.gpointer)(userData)

	retC := C.g_main_context_find_source_by_user_data((*C.GMainContext)(recv.native), c_user_data)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_main_context_get_poll_func : no return generator

// Runs a single iteration for the given main loop. This involves
// checking to see if any event sources are ready to be processed,
// then if no events sources are ready and @may_block is %TRUE, waiting
// for a source to become ready, then dispatching the highest priority
// events sources that are ready. Otherwise, if @may_block is %FALSE
// sources are not waited to become ready, only those highest priority
// events sources will be dispatched (if any), that are ready at this
// given moment without further waiting.
//
// Note that even when @may_block is %TRUE, it is still possible for
// g_main_context_iteration() to return %FALSE, since the wait may
// be interrupted for other reasons than an event source becoming ready.
/*

C function

g_main_context_iteration
*/
func (recv *MainContext) Iteration(mayBlock bool) bool {
	c_may_block :=
		boolToGboolean(mayBlock)

	retC := C.g_main_context_iteration((*C.GMainContext)(recv.native), c_may_block)
	retGo := retC == C.TRUE

	return retGo
}

// Checks if any sources have pending events for the given context.
/*

C function

g_main_context_pending
*/
func (recv *MainContext) Pending() bool {
	retC := C.g_main_context_pending((*C.GMainContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Prepares to poll sources within a main loop. The resulting information
// for polling is determined by calling g_main_context_query ().
//
// You must have successfully acquired the context with
// g_main_context_acquire() before you may call this function.
/*

C function

g_main_context_prepare
*/
func (recv *MainContext) Prepare(priority int32) bool {
	c_priority := (C.gint)(priority)

	retC := C.g_main_context_prepare((*C.GMainContext)(recv.native), &c_priority)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_main_context_query : unsupported parameter fds : output array param fds

// Increases the reference count on a #GMainContext object by one.
/*

C function

g_main_context_ref
*/
func (recv *MainContext) Ref() *MainContext {
	retC := C.g_main_context_ref((*C.GMainContext)(recv.native))
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Releases ownership of a context previously acquired by this thread
// with g_main_context_acquire(). If the context was acquired multiple
// times, the ownership will be released only when g_main_context_release()
// is called as many times as it was acquired.
/*

C function

g_main_context_release
*/
func (recv *MainContext) Release() {
	C.g_main_context_release((*C.GMainContext)(recv.native))

	return
}

// Removes file descriptor from the set of file descriptors to be
// polled for a particular context.
/*

C function

g_main_context_remove_poll
*/
func (recv *MainContext) RemovePoll(fd *PollFD) {
	c_fd := (*C.GPollFD)(C.NULL)
	if fd != nil {
		c_fd = (*C.GPollFD)(fd.ToC())
	}

	C.g_main_context_remove_poll((*C.GMainContext)(recv.native), c_fd)

	return
}

// Unsupported : g_main_context_set_poll_func : unsupported parameter func : no type generator for PollFunc (GPollFunc) for param func

// Decreases the reference count on a #GMainContext object by one. If
// the result is zero, free the context and free all associated memory.
/*

C function

g_main_context_unref
*/
func (recv *MainContext) Unref() {
	C.g_main_context_unref((*C.GMainContext)(recv.native))

	return
}

// Unsupported : g_main_context_wait : unsupported parameter mutex : no type generator for Mutex (GMutex*) for param mutex

// If @context is currently blocking in g_main_context_iteration()
// waiting for a source to become ready, cause it to stop blocking
// and return.  Otherwise, cause the next invocation of
// g_main_context_iteration() to return without blocking.
//
// This API is useful for low-level control over #GMainContext; for
// example, integrating it with main loop implementations such as
// #GMainLoop.
//
// Another related use for this function is when implementing a main
// loop with a termination condition, computed from multiple threads:
//
// |[<!-- language="C" -->
// #define NUM_TASKS 10
// static volatile gint tasks_remaining = NUM_TASKS;
// ...
//
// while (g_atomic_int_get (&tasks_remaining) != 0)
// g_main_context_iteration (NULL, TRUE);
// ]|
//
// Then in a thread:
// |[<!-- language="C" -->
// perform_work();
//
// if (g_atomic_int_dec_and_test (&tasks_remaining))
// g_main_context_wakeup (NULL);
// ]|
/*

C function

g_main_context_wakeup
*/
func (recv *MainContext) Wakeup() {
	C.g_main_context_wakeup((*C.GMainContext)(recv.native))

	return
}

// The `GMainLoop` struct is an opaque data type
// representing the main event loop of a GLib or GTK+ application.
/*

C type

GMainLoop
*/
type MainLoop struct {
	native *C.GMainLoop
}

func MainLoopNewFromC(u unsafe.Pointer) *MainLoop {
	c := (*C.GMainLoop)(u)
	if c == nil {
		return nil
	}

	g := &MainLoop{native: c}

	return g
}

func (recv *MainLoop) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Creates a new #GMainLoop structure.
/*

C function

g_main_loop_new
*/
func MainLoopNew(context *MainContext, isRunning bool) *MainLoop {
	c_context := (*C.GMainContext)(C.NULL)
	if context != nil {
		c_context = (*C.GMainContext)(context.ToC())
	}

	c_is_running :=
		boolToGboolean(isRunning)

	retC := C.g_main_loop_new(c_context, c_is_running)
	retGo := MainLoopNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns the #GMainContext of @loop.
/*

C function

g_main_loop_get_context
*/
func (recv *MainLoop) GetContext() *MainContext {
	retC := C.g_main_loop_get_context((*C.GMainLoop)(recv.native))
	retGo := MainContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Checks to see if the main loop is currently being run via g_main_loop_run().
/*

C function

g_main_loop_is_running
*/
func (recv *MainLoop) IsRunning() bool {
	retC := C.g_main_loop_is_running((*C.GMainLoop)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Stops a #GMainLoop from running. Any calls to g_main_loop_run()
// for the loop will return.
//
// Note that sources that have already been dispatched when
// g_main_loop_quit() is called will still be executed.
/*

C function

g_main_loop_quit
*/
func (recv *MainLoop) Quit() {
	C.g_main_loop_quit((*C.GMainLoop)(recv.native))

	return
}

// Increases the reference count on a #GMainLoop object by one.
/*

C function

g_main_loop_ref
*/
func (recv *MainLoop) Ref() *MainLoop {
	retC := C.g_main_loop_ref((*C.GMainLoop)(recv.native))
	retGo := MainLoopNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Runs a main loop until g_main_loop_quit() is called on the loop.
// If this is called for the thread of the loop's #GMainContext,
// it will process events from the loop, otherwise it will
// simply wait.
/*

C function

g_main_loop_run
*/
func (recv *MainLoop) Run() {
	C.g_main_loop_run((*C.GMainLoop)(recv.native))

	return
}

// Decreases the reference count on a #GMainLoop object by one. If
// the result is zero, free the loop and free all associated memory.
/*

C function

g_main_loop_unref
*/
func (recv *MainLoop) Unref() {
	C.g_main_loop_unref((*C.GMainLoop)(recv.native))

	return
}

// The #GMappedFile represents a file mapping created with
// g_mapped_file_new(). It has only private members and should
// not be accessed directly.
/*

C type

GMappedFile
*/
type MappedFile struct {
	native *C.GMappedFile
}

func MappedFileNewFromC(u unsafe.Pointer) *MappedFile {
	c := (*C.GMappedFile)(u)
	if c == nil {
		return nil
	}

	g := &MappedFile{native: c}

	return g
}

func (recv *MappedFile) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Decrements the reference count of @file by one.  If the reference count
// drops to 0, unmaps the buffer of @file and frees it.
//
// It is safe to call this function from any thread.
//
// Since 2.22
/*

C function

g_mapped_file_unref
*/
func (recv *MappedFile) Unref() {
	C.g_mapped_file_unref((*C.GMappedFile)(recv.native))

	return
}

// A parse context is used to parse a stream of bytes that
// you expect to contain marked-up text.
//
// See g_markup_parse_context_new(), #GMarkupParser, and so
// on for more details.
/*

C type

GMarkupParseContext
*/
type MarkupParseContext struct {
	native *C.GMarkupParseContext
}

func MarkupParseContextNewFromC(u unsafe.Pointer) *MarkupParseContext {
	c := (*C.GMarkupParseContext)(u)
	if c == nil {
		return nil
	}

	g := &MarkupParseContext{native: c}

	return g
}

func (recv *MarkupParseContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_markup_parse_context_new : unsupported parameter user_data_dnotify : no type generator for DestroyNotify (GDestroyNotify) for param user_data_dnotify

// Signals to the #GMarkupParseContext that all data has been
// fed into the parse context with g_markup_parse_context_parse().
//
// This function reports an error if the document isn't complete,
// for example if elements are still open.
/*

C function

g_markup_parse_context_end_parse
*/
func (recv *MarkupParseContext) EndParse() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_markup_parse_context_end_parse((*C.GMarkupParseContext)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Frees a #GMarkupParseContext.
//
// This function can't be called from inside one of the
// #GMarkupParser functions or while a subparser is pushed.
/*

C function

g_markup_parse_context_free
*/
func (recv *MarkupParseContext) Free() {
	C.g_markup_parse_context_free((*C.GMarkupParseContext)(recv.native))

	return
}

// Retrieves the current line number and the number of the character on
// that line. Intended for use in error messages; there are no strict
// semantics for what constitutes the "current" line number other than
// "the best number we could come up with for error messages."
/*

C function

g_markup_parse_context_get_position
*/
func (recv *MarkupParseContext) GetPosition(lineNumber int32, charNumber int32) {
	c_line_number := (C.gint)(lineNumber)

	c_char_number := (C.gint)(charNumber)

	C.g_markup_parse_context_get_position((*C.GMarkupParseContext)(recv.native), &c_line_number, &c_char_number)

	return
}

// Feed some data to the #GMarkupParseContext.
//
// The data need not be valid UTF-8; an error will be signaled if
// it's invalid. The data need not be an entire document; you can
// feed a document into the parser incrementally, via multiple calls
// to this function. Typically, as you receive data from a network
// connection or file, you feed each received chunk of data into this
// function, aborting the process if an error occurs. Once an error
// is reported, no further data may be fed to the #GMarkupParseContext;
// all errors are fatal.
/*

C function

g_markup_parse_context_parse
*/
func (recv *MarkupParseContext) Parse(text string, textLen int64) (bool, error) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_text_len := (C.gssize)(textLen)

	var cThrowableError *C.GError

	retC := C.g_markup_parse_context_parse((*C.GMarkupParseContext)(recv.native), c_text, c_text_len, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Any of the fields in #GMarkupParser can be %NULL, in which case they
// will be ignored. Except for the @error function, any of these callbacks
// can set an error; in particular the %G_MARKUP_ERROR_UNKNOWN_ELEMENT,
// %G_MARKUP_ERROR_UNKNOWN_ATTRIBUTE, and %G_MARKUP_ERROR_INVALID_CONTENT
// errors are intended to be set from these callbacks. If you set an error
// from a callback, g_markup_parse_context_parse() will report that error
// back to its caller.
/*

C type

GMarkupParser
*/
type MarkupParser struct {
	native *C.GMarkupParser
	// no type for start_element
	// no type for end_element
	// no type for text
	// no type for passthrough
	// no type for error
}

func MarkupParserNewFromC(u unsafe.Pointer) *MarkupParser {
	c := (*C.GMarkupParser)(u)
	if c == nil {
		return nil
	}

	g := &MarkupParser{native: c}

	return g
}

func (recv *MarkupParser) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// A GMatchInfo is an opaque struct used to return information about
// matches.
/*

C type

GMatchInfo
*/
type MatchInfo struct {
	native *C.GMatchInfo
}

func MatchInfoNewFromC(u unsafe.Pointer) *MatchInfo {
	c := (*C.GMatchInfo)(u)
	if c == nil {
		return nil
	}

	g := &MatchInfo{native: c}

	return g
}

func (recv *MatchInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// A set of functions used to perform memory allocation. The same #GMemVTable must
// be used for all allocations in the same program; a call to g_mem_set_vtable(),
// if it exists, should be prior to any use of GLib.
//
// This functions related to this has been deprecated in 2.46, and no longer work.
/*

C type

GMemVTable
*/
type MemVTable struct {
	native *C.GMemVTable
	// no type for malloc
	// no type for realloc
	// no type for free
	// no type for calloc
	// no type for try_malloc
	// no type for try_realloc
}

func MemVTableNewFromC(u unsafe.Pointer) *MemVTable {
	c := (*C.GMemVTable)(u)
	if c == nil {
		return nil
	}

	g := &MemVTable{native: c}

	return g
}

func (recv *MemVTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// The #GNode struct represents one node in a [n-ary tree][glib-N-ary-Trees].
/*

C type

GNode
*/
type Node struct {
	native *C.GNode
	Data   uintptr
	// next : record
	// prev : record
	// parent : record
	// children : record
}

func NodeNewFromC(u unsafe.Pointer) *Node {
	c := (*C.GNode)(u)
	if c == nil {
		return nil
	}

	g := &Node{
		Data:   (uintptr)(c.data),
		native: c,
	}

	return g
}

func (recv *Node) ToC() unsafe.Pointer {
	recv.native.data =
		(C.gpointer)(recv.Data)

	return (unsafe.Pointer)(recv.native)
}

// Gets the position of the first child of a #GNode
// which contains the given data.
/*

C function

g_node_child_index
*/
func (recv *Node) ChildIndex(data uintptr) int32 {
	c_data := (C.gpointer)(data)

	retC := C.g_node_child_index((*C.GNode)(recv.native), c_data)
	retGo := (int32)(retC)

	return retGo
}

// Gets the position of a #GNode with respect to its siblings.
// @child must be a child of @node. The first child is numbered 0,
// the second 1, and so on.
/*

C function

g_node_child_position
*/
func (recv *Node) ChildPosition(child *Node) int32 {
	c_child := (*C.GNode)(C.NULL)
	if child != nil {
		c_child = (*C.GNode)(child.ToC())
	}

	retC := C.g_node_child_position((*C.GNode)(recv.native), c_child)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_node_children_foreach : unsupported parameter func : no type generator for NodeForeachFunc (GNodeForeachFunc) for param func

// Recursively copies a #GNode (but does not deep-copy the data inside the
// nodes, see g_node_copy_deep() if you need that).
/*

C function

g_node_copy
*/
func (recv *Node) Copy() *Node {
	retC := C.g_node_copy((*C.GNode)(recv.native))
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the depth of a #GNode.
//
// If @node is %NULL the depth is 0. The root node has a depth of 1.
// For the children of the root node the depth is 2. And so on.
/*

C function

g_node_depth
*/
func (recv *Node) Depth() uint32 {
	retC := C.g_node_depth((*C.GNode)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Removes @root and its children from the tree, freeing any memory
// allocated.
/*

C function

g_node_destroy
*/
func (recv *Node) Destroy() {
	C.g_node_destroy((*C.GNode)(recv.native))

	return
}

// Finds a #GNode in a tree.
/*

C function

g_node_find
*/
func (recv *Node) Find(order TraverseType, flags TraverseFlags, data uintptr) *Node {
	c_order := (C.GTraverseType)(order)

	c_flags := (C.GTraverseFlags)(flags)

	c_data := (C.gpointer)(data)

	retC := C.g_node_find((*C.GNode)(recv.native), c_order, c_flags, c_data)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Finds the first child of a #GNode with the given data.
/*

C function

g_node_find_child
*/
func (recv *Node) FindChild(flags TraverseFlags, data uintptr) *Node {
	c_flags := (C.GTraverseFlags)(flags)

	c_data := (C.gpointer)(data)

	retC := C.g_node_find_child((*C.GNode)(recv.native), c_flags, c_data)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the first sibling of a #GNode.
// This could possibly be the node itself.
/*

C function

g_node_first_sibling
*/
func (recv *Node) FirstSibling() *Node {
	retC := C.g_node_first_sibling((*C.GNode)(recv.native))
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the root of a tree.
/*

C function

g_node_get_root
*/
func (recv *Node) GetRoot() *Node {
	retC := C.g_node_get_root((*C.GNode)(recv.native))
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Inserts a #GNode beneath the parent at the given position.
/*

C function

g_node_insert
*/
func (recv *Node) Insert(position int32, node *Node) *Node {
	c_position := (C.gint)(position)

	c_node := (*C.GNode)(C.NULL)
	if node != nil {
		c_node = (*C.GNode)(node.ToC())
	}

	retC := C.g_node_insert((*C.GNode)(recv.native), c_position, c_node)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Inserts a #GNode beneath the parent after the given sibling.
/*

C function

g_node_insert_after
*/
func (recv *Node) InsertAfter(sibling *Node, node *Node) *Node {
	c_sibling := (*C.GNode)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GNode)(sibling.ToC())
	}

	c_node := (*C.GNode)(C.NULL)
	if node != nil {
		c_node = (*C.GNode)(node.ToC())
	}

	retC := C.g_node_insert_after((*C.GNode)(recv.native), c_sibling, c_node)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Inserts a #GNode beneath the parent before the given sibling.
/*

C function

g_node_insert_before
*/
func (recv *Node) InsertBefore(sibling *Node, node *Node) *Node {
	c_sibling := (*C.GNode)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GNode)(sibling.ToC())
	}

	c_node := (*C.GNode)(C.NULL)
	if node != nil {
		c_node = (*C.GNode)(node.ToC())
	}

	retC := C.g_node_insert_before((*C.GNode)(recv.native), c_sibling, c_node)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Returns %TRUE if @node is an ancestor of @descendant.
// This is true if node is the parent of @descendant,
// or if node is the grandparent of @descendant etc.
/*

C function

g_node_is_ancestor
*/
func (recv *Node) IsAncestor(descendant *Node) bool {
	c_descendant := (*C.GNode)(C.NULL)
	if descendant != nil {
		c_descendant = (*C.GNode)(descendant.ToC())
	}

	retC := C.g_node_is_ancestor((*C.GNode)(recv.native), c_descendant)
	retGo := retC == C.TRUE

	return retGo
}

// Gets the last child of a #GNode.
/*

C function

g_node_last_child
*/
func (recv *Node) LastChild() *Node {
	retC := C.g_node_last_child((*C.GNode)(recv.native))
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the last sibling of a #GNode.
// This could possibly be the node itself.
/*

C function

g_node_last_sibling
*/
func (recv *Node) LastSibling() *Node {
	retC := C.g_node_last_sibling((*C.GNode)(recv.native))
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Gets the maximum height of all branches beneath a #GNode.
// This is the maximum distance from the #GNode to all leaf nodes.
//
// If @root is %NULL, 0 is returned. If @root has no children,
// 1 is returned. If @root has children, 2 is returned. And so on.
/*

C function

g_node_max_height
*/
func (recv *Node) MaxHeight() uint32 {
	retC := C.g_node_max_height((*C.GNode)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Gets the number of children of a #GNode.
/*

C function

g_node_n_children
*/
func (recv *Node) NChildren() uint32 {
	retC := C.g_node_n_children((*C.GNode)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Gets the number of nodes in a tree.
/*

C function

g_node_n_nodes
*/
func (recv *Node) NNodes(flags TraverseFlags) uint32 {
	c_flags := (C.GTraverseFlags)(flags)

	retC := C.g_node_n_nodes((*C.GNode)(recv.native), c_flags)
	retGo := (uint32)(retC)

	return retGo
}

// Gets a child of a #GNode, using the given index.
// The first child is at index 0. If the index is
// too big, %NULL is returned.
/*

C function

g_node_nth_child
*/
func (recv *Node) NthChild(n uint32) *Node {
	c_n := (C.guint)(n)

	retC := C.g_node_nth_child((*C.GNode)(recv.native), c_n)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Inserts a #GNode as the first child of the given parent.
/*

C function

g_node_prepend
*/
func (recv *Node) Prepend(node *Node) *Node {
	c_node := (*C.GNode)(C.NULL)
	if node != nil {
		c_node = (*C.GNode)(node.ToC())
	}

	retC := C.g_node_prepend((*C.GNode)(recv.native), c_node)
	retGo := NodeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Reverses the order of the children of a #GNode.
// (It doesn't change the order of the grandchildren.)
/*

C function

g_node_reverse_children
*/
func (recv *Node) ReverseChildren() {
	C.g_node_reverse_children((*C.GNode)(recv.native))

	return
}

// Unsupported : g_node_traverse : unsupported parameter func : no type generator for NodeTraverseFunc (GNodeTraverseFunc) for param func

// Unlinks a #GNode from a tree, resulting in two separate trees.
/*

C function

g_node_unlink
*/
func (recv *Node) Unlink() {
	C.g_node_unlink((*C.GNode)(recv.native))

	return
}

// A `GOptionContext` struct defines which options
// are accepted by the commandline option parser. The struct has only private
// fields and should not be directly accessed.
/*

C type

GOptionContext
*/
type OptionContext struct {
	native *C.GOptionContext
}

func OptionContextNewFromC(u unsafe.Pointer) *OptionContext {
	c := (*C.GOptionContext)(u)
	if c == nil {
		return nil
	}

	g := &OptionContext{native: c}

	return g
}

func (recv *OptionContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// A GOptionEntry struct defines a single option. To have an effect, they
// must be added to a #GOptionGroup with g_option_context_add_main_entries()
// or g_option_group_add_entries().
/*

C type

GOptionEntry
*/
type OptionEntry struct {
	native         *C.GOptionEntry
	LongName       string
	ShortName      rune
	Flags          int32
	Arg            OptionArg
	ArgData        uintptr
	Description    string
	ArgDescription string
}

func OptionEntryNewFromC(u unsafe.Pointer) *OptionEntry {
	c := (*C.GOptionEntry)(u)
	if c == nil {
		return nil
	}

	g := &OptionEntry{
		Arg:            (OptionArg)(c.arg),
		ArgData:        (uintptr)(c.arg_data),
		ArgDescription: C.GoString(c.arg_description),
		Description:    C.GoString(c.description),
		Flags:          (int32)(c.flags),
		LongName:       C.GoString(c.long_name),
		ShortName:      (rune)(c.short_name),
		native:         c,
	}

	return g
}

func (recv *OptionEntry) ToC() unsafe.Pointer {
	recv.native.long_name =
		C.CString(recv.LongName)
	recv.native.short_name =
		(C.gchar)(recv.ShortName)
	recv.native.flags =
		(C.gint)(recv.Flags)
	recv.native.arg =
		(C.GOptionArg)(recv.Arg)
	recv.native.arg_data =
		(C.gpointer)(recv.ArgData)
	recv.native.description =
		C.CString(recv.Description)
	recv.native.arg_description =
		C.CString(recv.ArgDescription)

	return (unsafe.Pointer)(recv.native)
}

// A `GOptionGroup` struct defines the options in a single
// group. The struct has only private fields and should not be directly accessed.
//
// All options in a group share the same translation function. Libraries which
// need to parse commandline options are expected to provide a function for
// getting a `GOptionGroup` holding their options, which
// the application can then add to its #GOptionContext.
/*

C type

GOptionGroup
*/
type OptionGroup struct {
	native *C.GOptionGroup
}

func OptionGroupNewFromC(u unsafe.Pointer) *OptionGroup {
	c := (*C.GOptionGroup)(u)
	if c == nil {
		return nil
	}

	g := &OptionGroup{native: c}

	return g
}

func (recv *OptionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// A GPatternSpec struct is the 'compiled' form of a pattern. This
// structure is opaque and its fields cannot be accessed directly.
/*

C type

GPatternSpec
*/
type PatternSpec struct {
	native *C.GPatternSpec
}

func PatternSpecNewFromC(u unsafe.Pointer) *PatternSpec {
	c := (*C.GPatternSpec)(u)
	if c == nil {
		return nil
	}

	g := &PatternSpec{native: c}

	return g
}

func (recv *PatternSpec) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Compares two compiled pattern specs and returns whether they will
// match the same set of strings.
/*

C function

g_pattern_spec_equal
*/
func (recv *PatternSpec) Equal(pspec2 *PatternSpec) bool {
	c_pspec2 := (*C.GPatternSpec)(C.NULL)
	if pspec2 != nil {
		c_pspec2 = (*C.GPatternSpec)(pspec2.ToC())
	}

	retC := C.g_pattern_spec_equal((*C.GPatternSpec)(recv.native), c_pspec2)
	retGo := retC == C.TRUE

	return retGo
}

// Frees the memory allocated for the #GPatternSpec.
/*

C function

g_pattern_spec_free
*/
func (recv *PatternSpec) Free() {
	C.g_pattern_spec_free((*C.GPatternSpec)(recv.native))

	return
}

// Represents a file descriptor, which events to poll for, and which events
// occurred.
/*

C type

GPollFD
*/
type PollFD struct {
	native  *C.GPollFD
	Fd      int32
	Events  uint32
	Revents uint32
}

func PollFDNewFromC(u unsafe.Pointer) *PollFD {
	c := (*C.GPollFD)(u)
	if c == nil {
		return nil
	}

	g := &PollFD{
		Events:  (uint32)(c.events),
		Fd:      (int32)(c.fd),
		Revents: (uint32)(c.revents),
		native:  c,
	}

	return g
}

func (recv *PollFD) ToC() unsafe.Pointer {
	recv.native.fd =
		(C.gint)(recv.Fd)
	recv.native.events =
		(C.gushort)(recv.Events)
	recv.native.revents =
		(C.gushort)(recv.Revents)

	return (unsafe.Pointer)(recv.native)
}

// The #GPrivate struct is an opaque data structure to represent a
// thread-local data key. It is approximately equivalent to the
// pthread_setspecific()/pthread_getspecific() APIs on POSIX and to
// TlsSetValue()/TlsGetValue() on Windows.
//
// If you don't already know why you might want this functionality,
// then you probably don't need it.
//
// #GPrivate is a very limited resource (as far as 128 per program,
// shared between all libraries). It is also not possible to destroy a
// #GPrivate after it has been used. As such, it is only ever acceptable
// to use #GPrivate in static scope, and even then sparingly so.
//
// See G_PRIVATE_INIT() for a couple of examples.
//
// The #GPrivate structure should be considered opaque.  It should only
// be accessed via the g_private_ functions.
/*

C type

GPrivate
*/
type Private struct {
	native *C.GPrivate
	// Private : p
	// Private : notify
	// Private : future
}

func PrivateNewFromC(u unsafe.Pointer) *Private {
	c := (*C.GPrivate)(u)
	if c == nil {
		return nil
	}

	g := &Private{native: c}

	return g
}

func (recv *Private) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Returns the current value of the thread local variable @key.
//
// If the value has not yet been set in this thread, %NULL is returned.
// Values are never copied between threads (when a new thread is
// created, for example).
/*

C function

g_private_get
*/
func (recv *Private) Get() uintptr {
	retC := C.g_private_get((*C.GPrivate)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Sets the thread local variable @key to have the value @value in the
// current thread.
//
// This function differs from g_private_replace() in the following way:
// the #GDestroyNotify for @key is not called on the old value.
/*

C function

g_private_set
*/
func (recv *Private) Set(value uintptr) {
	c_value := (C.gpointer)(value)

	C.g_private_set((*C.GPrivate)(recv.native), c_value)

	return
}

// Blacklisted : GPtrArray

// Contains the public fields of a
// [Queue][glib-Double-ended-Queues].
/*

C type

GQueue
*/
type Queue struct {
	native *C.GQueue
	// head : record
	// tail : record
	Length uint32
}

func QueueNewFromC(u unsafe.Pointer) *Queue {
	c := (*C.GQueue)(u)
	if c == nil {
		return nil
	}

	g := &Queue{
		Length: (uint32)(c.length),
		native: c,
	}

	return g
}

func (recv *Queue) ToC() unsafe.Pointer {
	recv.native.length =
		(C.guint)(recv.Length)

	return (unsafe.Pointer)(recv.native)
}

// Frees the memory allocated for the #GQueue. Only call this function
// if @queue was created with g_queue_new(). If queue elements contain
// dynamically-allocated memory, they should be freed first.
//
// If queue elements contain dynamically-allocated memory, you should
// either use g_queue_free_full() or free them manually first.
/*

C function

g_queue_free
*/
func (recv *Queue) Free() {
	C.g_queue_free((*C.GQueue)(recv.native))

	return
}

// Returns %TRUE if the queue is empty.
/*

C function

g_queue_is_empty
*/
func (recv *Queue) IsEmpty() bool {
	retC := C.g_queue_is_empty((*C.GQueue)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Returns the first element of the queue.
/*

C function

g_queue_peek_head
*/
func (recv *Queue) PeekHead() uintptr {
	retC := C.g_queue_peek_head((*C.GQueue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Returns the last element of the queue.
/*

C function

g_queue_peek_tail
*/
func (recv *Queue) PeekTail() uintptr {
	retC := C.g_queue_peek_tail((*C.GQueue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Removes the first element of the queue and returns its data.
/*

C function

g_queue_pop_head
*/
func (recv *Queue) PopHead() uintptr {
	retC := C.g_queue_pop_head((*C.GQueue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Removes and returns the first element of the queue.
/*

C function

g_queue_pop_head_link
*/
func (recv *Queue) PopHeadLink() *List {
	retC := C.g_queue_pop_head_link((*C.GQueue)(recv.native))
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Removes the last element of the queue and returns its data.
/*

C function

g_queue_pop_tail
*/
func (recv *Queue) PopTail() uintptr {
	retC := C.g_queue_pop_tail((*C.GQueue)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Removes and returns the last element of the queue.
/*

C function

g_queue_pop_tail_link
*/
func (recv *Queue) PopTailLink() *List {
	retC := C.g_queue_pop_tail_link((*C.GQueue)(recv.native))
	retGo := ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds a new element at the head of the queue.
/*

C function

g_queue_push_head
*/
func (recv *Queue) PushHead(data uintptr) {
	c_data := (C.gpointer)(data)

	C.g_queue_push_head((*C.GQueue)(recv.native), c_data)

	return
}

// Adds a new element at the head of the queue.
/*

C function

g_queue_push_head_link
*/
func (recv *Queue) PushHeadLink(link *List) {
	c_link_ := (*C.GList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GList)(link.ToC())
	}

	C.g_queue_push_head_link((*C.GQueue)(recv.native), c_link_)

	return
}

// Adds a new element at the tail of the queue.
/*

C function

g_queue_push_tail
*/
func (recv *Queue) PushTail(data uintptr) {
	c_data := (C.gpointer)(data)

	C.g_queue_push_tail((*C.GQueue)(recv.native), c_data)

	return
}

// Adds a new element at the tail of the queue.
/*

C function

g_queue_push_tail_link
*/
func (recv *Queue) PushTailLink(link *List) {
	c_link_ := (*C.GList)(C.NULL)
	if link != nil {
		c_link_ = (*C.GList)(link.ToC())
	}

	C.g_queue_push_tail_link((*C.GQueue)(recv.native), c_link_)

	return
}

// The GRand struct is an opaque data structure. It should only be
// accessed through the g_rand_* functions.
/*

C type

GRand
*/
type Rand struct {
	native *C.GRand
}

func RandNewFromC(u unsafe.Pointer) *Rand {
	c := (*C.GRand)(u)
	if c == nil {
		return nil
	}

	g := &Rand{native: c}

	return g
}

func (recv *Rand) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Returns the next random #gdouble from @rand_ equally distributed over
// the range [0..1).
/*

C function

g_rand_double
*/
func (recv *Rand) Double() float64 {
	retC := C.g_rand_double((*C.GRand)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// Returns the next random #gdouble from @rand_ equally distributed over
// the range [@begin..@end).
/*

C function

g_rand_double_range
*/
func (recv *Rand) DoubleRange(begin float64, end float64) float64 {
	c_begin := (C.gdouble)(begin)

	c_end := (C.gdouble)(end)

	retC := C.g_rand_double_range((*C.GRand)(recv.native), c_begin, c_end)
	retGo := (float64)(retC)

	return retGo
}

// Frees the memory allocated for the #GRand.
/*

C function

g_rand_free
*/
func (recv *Rand) Free() {
	C.g_rand_free((*C.GRand)(recv.native))

	return
}

// Returns the next random #guint32 from @rand_ equally distributed over
// the range [0..2^32-1].
/*

C function

g_rand_int
*/
func (recv *Rand) Int() uint32 {
	retC := C.g_rand_int((*C.GRand)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Returns the next random #gint32 from @rand_ equally distributed over
// the range [@begin..@end-1].
/*

C function

g_rand_int_range
*/
func (recv *Rand) IntRange(begin int32, end int32) int32 {
	c_begin := (C.gint32)(begin)

	c_end := (C.gint32)(end)

	retC := C.g_rand_int_range((*C.GRand)(recv.native), c_begin, c_end)
	retGo := (int32)(retC)

	return retGo
}

// Sets the seed for the random number generator #GRand to @seed.
/*

C function

g_rand_set_seed
*/
func (recv *Rand) SetSeed(seed uint32) {
	c_seed := (C.guint32)(seed)

	C.g_rand_set_seed((*C.GRand)(recv.native), c_seed)

	return
}

// The #GSList struct is used for each element in the singly-linked
// list.
/*

C type

GSList
*/
type SList struct {
	native *C.GSList
	Data   uintptr
	// next : record
}

func SListNewFromC(u unsafe.Pointer) *SList {
	c := (*C.GSList)(u)
	if c == nil {
		return nil
	}

	g := &SList{
		Data:   (uintptr)(c.data),
		native: c,
	}

	return g
}

func (recv *SList) ToC() unsafe.Pointer {
	recv.native.data =
		(C.gpointer)(recv.Data)

	return (unsafe.Pointer)(recv.native)
}

// The data structure representing a lexical scanner.
//
// You should set @input_name after creating the scanner, since
// it is used by the default message handler when displaying
// warnings and errors. If you are scanning a file, the filename
// would be a good choice.
//
// The @user_data and @max_parse_errors fields are not used.
// If you need to associate extra data with the scanner you
// can place them here.
//
// If you want to use your own message handler you can set the
// @msg_handler field. The type of the message handler function
// is declared by #GScannerMsgFunc.
/*

C type

GScanner
*/
type Scanner struct {
	native         *C.GScanner
	UserData       uintptr
	MaxParseErrors uint32
	ParseErrors    uint32
	InputName      string
	// qdata : record
	// config : record
	Token TokenType
	// value : no type generator for TokenValue, GTokenValue
	Line      uint32
	Position  uint32
	NextToken TokenType
	// next_value : no type generator for TokenValue, GTokenValue
	NextLine     uint32
	NextPosition uint32
	// Private : symbol_table
	// Private : input_fd
	// Private : text
	// Private : text_end
	// Private : buffer
	// Private : scope_id
	// msg_handler : no type generator for ScannerMsgFunc, GScannerMsgFunc
}

func ScannerNewFromC(u unsafe.Pointer) *Scanner {
	c := (*C.GScanner)(u)
	if c == nil {
		return nil
	}

	g := &Scanner{
		InputName:      C.GoString(c.input_name),
		Line:           (uint32)(c.line),
		MaxParseErrors: (uint32)(c.max_parse_errors),
		NextLine:       (uint32)(c.next_line),
		NextPosition:   (uint32)(c.next_position),
		NextToken:      (TokenType)(c.next_token),
		ParseErrors:    (uint32)(c.parse_errors),
		Position:       (uint32)(c.position),
		Token:          (TokenType)(c.token),
		UserData:       (uintptr)(c.user_data),
		native:         c,
	}

	return g
}

func (recv *Scanner) ToC() unsafe.Pointer {
	recv.native.user_data =
		(C.gpointer)(recv.UserData)
	recv.native.max_parse_errors =
		(C.guint)(recv.MaxParseErrors)
	recv.native.parse_errors =
		(C.guint)(recv.ParseErrors)
	recv.native.input_name =
		C.CString(recv.InputName)
	recv.native.token =
		(C.GTokenType)(recv.Token)
	recv.native.line =
		(C.guint)(recv.Line)
	recv.native.position =
		(C.guint)(recv.Position)
	recv.native.next_token =
		(C.GTokenType)(recv.NextToken)
	recv.native.next_line =
		(C.guint)(recv.NextLine)
	recv.native.next_position =
		(C.guint)(recv.NextPosition)

	return (unsafe.Pointer)(recv.native)
}

// Returns the current line in the input stream (counting
// from 1). This is the line of the last token parsed via
// g_scanner_get_next_token().
/*

C function

g_scanner_cur_line
*/
func (recv *Scanner) CurLine() uint32 {
	retC := C.g_scanner_cur_line((*C.GScanner)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Returns the current position in the current line (counting
// from 0). This is the position of the last token parsed via
// g_scanner_get_next_token().
/*

C function

g_scanner_cur_position
*/
func (recv *Scanner) CurPosition() uint32 {
	retC := C.g_scanner_cur_position((*C.GScanner)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Gets the current token type. This is simply the @token
// field in the #GScanner structure.
/*

C function

g_scanner_cur_token
*/
func (recv *Scanner) CurToken() TokenType {
	retC := C.g_scanner_cur_token((*C.GScanner)(recv.native))
	retGo := (TokenType)(retC)

	return retGo
}

// Unsupported : g_scanner_cur_value : no return generator

// Frees all memory used by the #GScanner.
/*

C function

g_scanner_destroy
*/
func (recv *Scanner) Destroy() {
	C.g_scanner_destroy((*C.GScanner)(recv.native))

	return
}

// Returns %TRUE if the scanner has reached the end of
// the file or text buffer.
/*

C function

g_scanner_eof
*/
func (recv *Scanner) Eof() bool {
	retC := C.g_scanner_eof((*C.GScanner)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_scanner_error : unsupported parameter ... : varargs

// Parses the next token just like g_scanner_peek_next_token()
// and also removes it from the input stream. The token data is
// placed in the @token, @value, @line, and @position fields of
// the #GScanner structure.
/*

C function

g_scanner_get_next_token
*/
func (recv *Scanner) GetNextToken() TokenType {
	retC := C.g_scanner_get_next_token((*C.GScanner)(recv.native))
	retGo := (TokenType)(retC)

	return retGo
}

// Prepares to scan a file.
/*

C function

g_scanner_input_file
*/
func (recv *Scanner) InputFile(inputFd int32) {
	c_input_fd := (C.gint)(inputFd)

	C.g_scanner_input_file((*C.GScanner)(recv.native), c_input_fd)

	return
}

// Prepares to scan a text buffer.
/*

C function

g_scanner_input_text
*/
func (recv *Scanner) InputText(text string, textLen uint32) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_text_len := (C.guint)(textLen)

	C.g_scanner_input_text((*C.GScanner)(recv.native), c_text, c_text_len)

	return
}

// Looks up a symbol in the current scope and return its value.
// If the symbol is not bound in the current scope, %NULL is
// returned.
/*

C function

g_scanner_lookup_symbol
*/
func (recv *Scanner) LookupSymbol(symbol string) uintptr {
	c_symbol := C.CString(symbol)
	defer C.free(unsafe.Pointer(c_symbol))

	retC := C.g_scanner_lookup_symbol((*C.GScanner)(recv.native), c_symbol)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Parses the next token, without removing it from the input stream.
// The token data is placed in the @next_token, @next_value, @next_line,
// and @next_position fields of the #GScanner structure.
//
// Note that, while the token is not removed from the input stream
// (i.e. the next call to g_scanner_get_next_token() will return the
// same token), it will not be reevaluated. This can lead to surprising
// results when changing scope or the scanner configuration after peeking
// the next token. Getting the next token after switching the scope or
// configuration will return whatever was peeked before, regardless of
// any symbols that may have been added or removed in the new scope.
/*

C function

g_scanner_peek_next_token
*/
func (recv *Scanner) PeekNextToken() TokenType {
	retC := C.g_scanner_peek_next_token((*C.GScanner)(recv.native))
	retGo := (TokenType)(retC)

	return retGo
}

// Adds a symbol to the given scope.
/*

C function

g_scanner_scope_add_symbol
*/
func (recv *Scanner) ScopeAddSymbol(scopeId uint32, symbol string, value uintptr) {
	c_scope_id := (C.guint)(scopeId)

	c_symbol := C.CString(symbol)
	defer C.free(unsafe.Pointer(c_symbol))

	c_value := (C.gpointer)(value)

	C.g_scanner_scope_add_symbol((*C.GScanner)(recv.native), c_scope_id, c_symbol, c_value)

	return
}

// Unsupported : g_scanner_scope_foreach_symbol : unsupported parameter func : no type generator for HFunc (GHFunc) for param func

// Looks up a symbol in a scope and return its value. If the
// symbol is not bound in the scope, %NULL is returned.
/*

C function

g_scanner_scope_lookup_symbol
*/
func (recv *Scanner) ScopeLookupSymbol(scopeId uint32, symbol string) uintptr {
	c_scope_id := (C.guint)(scopeId)

	c_symbol := C.CString(symbol)
	defer C.free(unsafe.Pointer(c_symbol))

	retC := C.g_scanner_scope_lookup_symbol((*C.GScanner)(recv.native), c_scope_id, c_symbol)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Removes a symbol from a scope.
/*

C function

g_scanner_scope_remove_symbol
*/
func (recv *Scanner) ScopeRemoveSymbol(scopeId uint32, symbol string) {
	c_scope_id := (C.guint)(scopeId)

	c_symbol := C.CString(symbol)
	defer C.free(unsafe.Pointer(c_symbol))

	C.g_scanner_scope_remove_symbol((*C.GScanner)(recv.native), c_scope_id, c_symbol)

	return
}

// Sets the current scope.
/*

C function

g_scanner_set_scope
*/
func (recv *Scanner) SetScope(scopeId uint32) uint32 {
	c_scope_id := (C.guint)(scopeId)

	retC := C.g_scanner_set_scope((*C.GScanner)(recv.native), c_scope_id)
	retGo := (uint32)(retC)

	return retGo
}

// Rewinds the filedescriptor to the current buffer position
// and blows the file read ahead buffer. This is useful for
// third party uses of the scanners filedescriptor, which hooks
// onto the current scanning position.
/*

C function

g_scanner_sync_file_offset
*/
func (recv *Scanner) SyncFileOffset() {
	C.g_scanner_sync_file_offset((*C.GScanner)(recv.native))

	return
}

// Outputs a message through the scanner's msg_handler,
// resulting from an unexpected token in the input stream.
// Note that you should not call g_scanner_peek_next_token()
// followed by g_scanner_unexp_token() without an intermediate
// call to g_scanner_get_next_token(), as g_scanner_unexp_token()
// evaluates the scanner's current token (not the peeked token)
// to construct part of the message.
/*

C function

g_scanner_unexp_token
*/
func (recv *Scanner) UnexpToken(expectedToken TokenType, identifierSpec string, symbolSpec string, symbolName string, message string, isError int32) {
	c_expected_token := (C.GTokenType)(expectedToken)

	c_identifier_spec := C.CString(identifierSpec)
	defer C.free(unsafe.Pointer(c_identifier_spec))

	c_symbol_spec := C.CString(symbolSpec)
	defer C.free(unsafe.Pointer(c_symbol_spec))

	c_symbol_name := C.CString(symbolName)
	defer C.free(unsafe.Pointer(c_symbol_name))

	c_message := C.CString(message)
	defer C.free(unsafe.Pointer(c_message))

	c_is_error := (C.gint)(isError)

	C.g_scanner_unexp_token((*C.GScanner)(recv.native), c_expected_token, c_identifier_spec, c_symbol_spec, c_symbol_name, c_message, c_is_error)

	return
}

// Unsupported : g_scanner_warn : unsupported parameter ... : varargs

// Specifies the #GScanner parser configuration. Most settings can
// be changed during the parsing phase and will affect the lexical
// parsing of the next unpeeked token.
/*

C type

GScannerConfig
*/
type ScannerConfig struct {
	native              *C.GScannerConfig
	CsetSkipCharacters  string
	CsetIdentifierFirst string
	CsetIdentifierNth   string
	CpairCommentSingle  string
	// Bitfield not supported :  1 case_sensitive
	// Bitfield not supported :  1 skip_comment_multi
	// Bitfield not supported :  1 skip_comment_single
	// Bitfield not supported :  1 scan_comment_multi
	// Bitfield not supported :  1 scan_identifier
	// Bitfield not supported :  1 scan_identifier_1char
	// Bitfield not supported :  1 scan_identifier_NULL
	// Bitfield not supported :  1 scan_symbols
	// Bitfield not supported :  1 scan_binary
	// Bitfield not supported :  1 scan_octal
	// Bitfield not supported :  1 scan_float
	// Bitfield not supported :  1 scan_hex
	// Bitfield not supported :  1 scan_hex_dollar
	// Bitfield not supported :  1 scan_string_sq
	// Bitfield not supported :  1 scan_string_dq
	// Bitfield not supported :  1 numbers_2_int
	// Bitfield not supported :  1 int_2_float
	// Bitfield not supported :  1 identifier_2_string
	// Bitfield not supported :  1 char_2_token
	// Bitfield not supported :  1 symbol_2_token
	// Bitfield not supported :  1 scope_0_fallback
	// Bitfield not supported :  1 store_int64
	// Private : padding_dummy
}

func ScannerConfigNewFromC(u unsafe.Pointer) *ScannerConfig {
	c := (*C.GScannerConfig)(u)
	if c == nil {
		return nil
	}

	g := &ScannerConfig{
		CpairCommentSingle:  C.GoString(c.cpair_comment_single),
		CsetIdentifierFirst: C.GoString(c.cset_identifier_first),
		CsetIdentifierNth:   C.GoString(c.cset_identifier_nth),
		CsetSkipCharacters:  C.GoString(c.cset_skip_characters),
		native:              c,
	}

	return g
}

func (recv *ScannerConfig) ToC() unsafe.Pointer {
	recv.native.cset_skip_characters =
		C.CString(recv.CsetSkipCharacters)
	recv.native.cset_identifier_first =
		C.CString(recv.CsetIdentifierFirst)
	recv.native.cset_identifier_nth =
		C.CString(recv.CsetIdentifierNth)
	recv.native.cpair_comment_single =
		C.CString(recv.CpairCommentSingle)

	return (unsafe.Pointer)(recv.native)
}

// The #GSequence struct is an opaque data type representing a
// [sequence][glib-Sequences] data type.
/*

C type

GSequence
*/
type Sequence struct {
	native *C.GSequence
}

func SequenceNewFromC(u unsafe.Pointer) *Sequence {
	c := (*C.GSequence)(u)
	if c == nil {
		return nil
	}

	g := &Sequence{native: c}

	return g
}

func (recv *Sequence) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// The #GSequenceIter struct is an opaque data type representing an
// iterator pointing into a #GSequence.
/*

C type

GSequenceIter
*/
type SequenceIter struct {
	native *C.GSequenceIter
}

func SequenceIterNewFromC(u unsafe.Pointer) *SequenceIter {
	c := (*C.GSequenceIter)(u)
	if c == nil {
		return nil
	}

	g := &SequenceIter{native: c}

	return g
}

func (recv *SequenceIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// The `GSource` struct is an opaque data type
// representing an event source.
/*

C type

GSource
*/
type Source struct {
	native *C.GSource
	// Private : callback_data
	// Private : callback_funcs
	// Private : source_funcs
	// Private : ref_count
	// Private : context
	// Private : priority
	// Private : flags
	// Private : source_id
	// Private : poll_fds
	// Private : prev
	// Private : next
	// Private : name
	// Private : priv
}

func SourceNewFromC(u unsafe.Pointer) *Source {
	c := (*C.GSource)(u)
	if c == nil {
		return nil
	}

	g := &Source{native: c}

	return g
}

func (recv *Source) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Creates a new #GSource structure. The size is specified to
// allow creating structures derived from #GSource that contain
// additional data. The size passed in must be at least
// `sizeof (GSource)`.
//
// The source will not initially be associated with any #GMainContext
// and must be added to one with g_source_attach() before it will be
// executed.
/*

C function

g_source_new
*/
func SourceNew(sourceFuncs *SourceFuncs, structSize uint32) *Source {
	c_source_funcs := (*C.GSourceFuncs)(C.NULL)
	if sourceFuncs != nil {
		c_source_funcs = (*C.GSourceFuncs)(sourceFuncs.ToC())
	}

	c_struct_size := (C.guint)(structSize)

	retC := C.g_source_new(c_source_funcs, c_struct_size)
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds a file descriptor to the set of file descriptors polled for
// this source. This is usually combined with g_source_new() to add an
// event source. The event source's check function will typically test
// the @revents field in the #GPollFD struct and return %TRUE if events need
// to be processed.
//
// This API is only intended to be used by implementations of #GSource.
// Do not call this API on a #GSource that you did not create.
//
// Using this API forces the linear scanning of event sources on each
// main loop iteration.  Newly-written event sources should try to use
// g_source_add_unix_fd() instead of this API.
/*

C function

g_source_add_poll
*/
func (recv *Source) AddPoll(fd *PollFD) {
	c_fd := (*C.GPollFD)(C.NULL)
	if fd != nil {
		c_fd = (*C.GPollFD)(fd.ToC())
	}

	C.g_source_add_poll((*C.GSource)(recv.native), c_fd)

	return
}

// Adds a #GSource to a @context so that it will be executed within
// that context. Remove it by calling g_source_destroy().
/*

C function

g_source_attach
*/
func (recv *Source) Attach(context *MainContext) uint32 {
	c_context := (*C.GMainContext)(C.NULL)
	if context != nil {
		c_context = (*C.GMainContext)(context.ToC())
	}

	retC := C.g_source_attach((*C.GSource)(recv.native), c_context)
	retGo := (uint32)(retC)

	return retGo
}

// Removes a source from its #GMainContext, if any, and mark it as
// destroyed.  The source cannot be subsequently added to another
// context. It is safe to call this on sources which have already been
// removed from their context.
/*

C function

g_source_destroy
*/
func (recv *Source) Destroy() {
	C.g_source_destroy((*C.GSource)(recv.native))

	return
}

// Checks whether a source is allowed to be called recursively.
// see g_source_set_can_recurse().
/*

C function

g_source_get_can_recurse
*/
func (recv *Source) GetCanRecurse() bool {
	retC := C.g_source_get_can_recurse((*C.GSource)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Gets the #GMainContext with which the source is associated.
//
// You can call this on a source that has been destroyed, provided
// that the #GMainContext it was attached to still exists (in which
// case it will return that #GMainContext). In particular, you can
// always call this function on the source returned from
// g_main_current_source(). But calling this function on a source
// whose #GMainContext has been destroyed is an error.
/*

C function

g_source_get_context
*/
func (recv *Source) GetContext() *MainContext {
	retC := C.g_source_get_context((*C.GSource)(recv.native))
	var retGo (*MainContext)
	if retC == nil {
		retGo = nil
	} else {
		retGo = MainContextNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// This function ignores @source and is otherwise the same as
// g_get_current_time().
/*

C function

g_source_get_current_time
*/
func (recv *Source) GetCurrentTime(timeval *TimeVal) {
	c_timeval := (*C.GTimeVal)(C.NULL)
	if timeval != nil {
		c_timeval = (*C.GTimeVal)(timeval.ToC())
	}

	C.g_source_get_current_time((*C.GSource)(recv.native), c_timeval)

	return
}

// Returns the numeric ID for a particular source. The ID of a source
// is a positive integer which is unique within a particular main loop
// context. The reverse
// mapping from ID to source is done by g_main_context_find_source_by_id().
/*

C function

g_source_get_id
*/
func (recv *Source) GetId() uint32 {
	retC := C.g_source_get_id((*C.GSource)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Gets the priority of a source.
/*

C function

g_source_get_priority
*/
func (recv *Source) GetPriority() int32 {
	retC := C.g_source_get_priority((*C.GSource)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Gets the "ready time" of @source, as set by
// g_source_set_ready_time().
//
// Any time before the current monotonic time (including 0) is an
// indication that the source will fire immediately.
/*

C function

g_source_get_ready_time
*/
func (recv *Source) GetReadyTime() int64 {
	retC := C.g_source_get_ready_time((*C.GSource)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Increases the reference count on a source by one.
/*

C function

g_source_ref
*/
func (recv *Source) Ref() *Source {
	retC := C.g_source_ref((*C.GSource)(recv.native))
	retGo := SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Removes a file descriptor from the set of file descriptors polled for
// this source.
//
// This API is only intended to be used by implementations of #GSource.
// Do not call this API on a #GSource that you did not create.
/*

C function

g_source_remove_poll
*/
func (recv *Source) RemovePoll(fd *PollFD) {
	c_fd := (*C.GPollFD)(C.NULL)
	if fd != nil {
		c_fd = (*C.GPollFD)(fd.ToC())
	}

	C.g_source_remove_poll((*C.GSource)(recv.native), c_fd)

	return
}

// Unsupported : g_source_set_callback : unsupported parameter func : no type generator for SourceFunc (GSourceFunc) for param func

// Sets the callback function storing the data as a refcounted callback
// "object". This is used internally. Note that calling
// g_source_set_callback_indirect() assumes
// an initial reference count on @callback_data, and thus
// @callback_funcs->unref will eventually be called once more
// than @callback_funcs->ref.
/*

C function

g_source_set_callback_indirect
*/
func (recv *Source) SetCallbackIndirect(callbackData uintptr, callbackFuncs *SourceCallbackFuncs) {
	c_callback_data := (C.gpointer)(callbackData)

	c_callback_funcs := (*C.GSourceCallbackFuncs)(C.NULL)
	if callbackFuncs != nil {
		c_callback_funcs = (*C.GSourceCallbackFuncs)(callbackFuncs.ToC())
	}

	C.g_source_set_callback_indirect((*C.GSource)(recv.native), c_callback_data, c_callback_funcs)

	return
}

// Sets whether a source can be called recursively. If @can_recurse is
// %TRUE, then while the source is being dispatched then this source
// will be processed normally. Otherwise, all processing of this
// source is blocked until the dispatch function returns.
/*

C function

g_source_set_can_recurse
*/
func (recv *Source) SetCanRecurse(canRecurse bool) {
	c_can_recurse :=
		boolToGboolean(canRecurse)

	C.g_source_set_can_recurse((*C.GSource)(recv.native), c_can_recurse)

	return
}

// Sets the priority of a source. While the main loop is being run, a
// source will be dispatched if it is ready to be dispatched and no
// sources at a higher (numerically smaller) priority are ready to be
// dispatched.
//
// A child source always has the same priority as its parent.  It is not
// permitted to change the priority of a source once it has been added
// as a child of another source.
/*

C function

g_source_set_priority
*/
func (recv *Source) SetPriority(priority int32) {
	c_priority := (C.gint)(priority)

	C.g_source_set_priority((*C.GSource)(recv.native), c_priority)

	return
}

// Decreases the reference count of a source by one. If the
// resulting reference count is zero the source and associated
// memory will be destroyed.
/*

C function

g_source_unref
*/
func (recv *Source) Unref() {
	C.g_source_unref((*C.GSource)(recv.native))

	return
}

// The `GSourceCallbackFuncs` struct contains
// functions for managing callback objects.
/*

C type

GSourceCallbackFuncs
*/
type SourceCallbackFuncs struct {
	native *C.GSourceCallbackFuncs
	// no type for ref
	// no type for unref
	// no type for get
}

func SourceCallbackFuncsNewFromC(u unsafe.Pointer) *SourceCallbackFuncs {
	c := (*C.GSourceCallbackFuncs)(u)
	if c == nil {
		return nil
	}

	g := &SourceCallbackFuncs{native: c}

	return g
}

func (recv *SourceCallbackFuncs) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// The `GSourceFuncs` struct contains a table of
// functions used to handle event sources in a generic manner.
//
// For idle sources, the prepare and check functions always return %TRUE
// to indicate that the source is always ready to be processed. The prepare
// function also returns a timeout value of 0 to ensure that the poll() call
// doesn't block (since that would be time wasted which could have been spent
// running the idle function).
//
// For timeout sources, the prepare and check functions both return %TRUE
// if the timeout interval has expired. The prepare function also returns
// a timeout value to ensure that the poll() call doesn't block too long
// and miss the next timeout.
//
// For file descriptor sources, the prepare function typically returns %FALSE,
// since it must wait until poll() has been called before it knows whether
// any events need to be processed. It sets the returned timeout to -1 to
// indicate that it doesn't mind how long the poll() call blocks. In the
// check function, it tests the results of the poll() call to see if the
// required condition has been met, and returns %TRUE if so.
/*

C type

GSourceFuncs
*/
type SourceFuncs struct {
	native *C.GSourceFuncs
	// no type for prepare
	// no type for check
	// no type for dispatch
	// no type for finalize
	// Private : closure_callback
	// Private : closure_marshal
}

func SourceFuncsNewFromC(u unsafe.Pointer) *SourceFuncs {
	c := (*C.GSourceFuncs)(u)
	if c == nil {
		return nil
	}

	g := &SourceFuncs{native: c}

	return g
}

func (recv *SourceFuncs) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GSourcePrivate
*/
type SourcePrivate struct {
	native *C.GSourcePrivate
}

func SourcePrivateNewFromC(u unsafe.Pointer) *SourcePrivate {
	c := (*C.GSourcePrivate)(u)
	if c == nil {
		return nil
	}

	g := &SourcePrivate{native: c}

	return g
}

func (recv *SourcePrivate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// A type corresponding to the appropriate struct type for the stat()
// system call, depending on the platform and/or compiler being used.
//
// See g_stat() for more information.
/*

C type

GStatBuf
*/
type StatBuf struct {
	native *C.GStatBuf
}

func StatBufNewFromC(u unsafe.Pointer) *StatBuf {
	c := (*C.GStatBuf)(u)
	if c == nil {
		return nil
	}

	g := &StatBuf{native: c}

	return g
}

func (recv *StatBuf) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// The GString struct contains the public fields of a GString.
/*

C type

GString
*/
type String struct {
	native       *C.GString
	Str          string
	Len          uint64
	AllocatedLen uint64
}

func StringNewFromC(u unsafe.Pointer) *String {
	c := (*C.GString)(u)
	if c == nil {
		return nil
	}

	g := &String{
		AllocatedLen: (uint64)(c.allocated_len),
		Len:          (uint64)(c.len),
		Str:          C.GoString(c.str),
		native:       c,
	}

	return g
}

func (recv *String) ToC() unsafe.Pointer {
	recv.native.str =
		C.CString(recv.Str)
	recv.native.len =
		(C.gsize)(recv.Len)
	recv.native.allocated_len =
		(C.gsize)(recv.AllocatedLen)

	return (unsafe.Pointer)(recv.native)
}

// Adds a string onto the end of a #GString, expanding
// it if necessary.
/*

C function

g_string_append
*/
func (recv *String) Append(val string) *String {
	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	retC := C.g_string_append((*C.GString)(recv.native), c_val)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds a byte onto the end of a #GString, expanding
// it if necessary.
/*

C function

g_string_append_c
*/
func (recv *String) AppendC(c rune) *String {
	c_c := (C.gchar)(c)

	retC := C.g_string_append_c((*C.GString)(recv.native), c_c)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Appends @len bytes of @val to @string. Because @len is
// provided, @val may contain embedded nuls and need not
// be nul-terminated.
//
// Since this function does not stop at nul bytes, it is
// the caller's responsibility to ensure that @val has at
// least @len addressable bytes.
/*

C function

g_string_append_len
*/
func (recv *String) AppendLen(val string, len int64) *String {
	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	c_len := (C.gssize)(len)

	retC := C.g_string_append_len((*C.GString)(recv.native), c_val, c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_string_append_printf : unsupported parameter ... : varargs

// Converts a Unicode character into UTF-8, and appends it
// to the string.
/*

C function

g_string_append_unichar
*/
func (recv *String) AppendUnichar(wc rune) *String {
	c_wc := (C.gunichar)(wc)

	retC := C.g_string_append_unichar((*C.GString)(recv.native), c_wc)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Converts all uppercase ASCII letters to lowercase ASCII letters.
/*

C function

g_string_ascii_down
*/
func (recv *String) AsciiDown() *String {
	retC := C.g_string_ascii_down((*C.GString)(recv.native))
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Converts all lowercase ASCII letters to uppercase ASCII letters.
/*

C function

g_string_ascii_up
*/
func (recv *String) AsciiUp() *String {
	retC := C.g_string_ascii_up((*C.GString)(recv.native))
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copies the bytes from a string into a #GString,
// destroying any previous contents. It is rather like
// the standard strcpy() function, except that you do not
// have to worry about having enough space to copy the string.
/*

C function

g_string_assign
*/
func (recv *String) Assign(rval string) *String {
	c_rval := C.CString(rval)
	defer C.free(unsafe.Pointer(c_rval))

	retC := C.g_string_assign((*C.GString)(recv.native), c_rval)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Converts a #GString to lowercase.
/*

C function

g_string_down
*/
func (recv *String) Down() *String {
	retC := C.g_string_down((*C.GString)(recv.native))
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Compares two strings for equality, returning %TRUE if they are equal.
// For use with #GHashTable.
/*

C function

g_string_equal
*/
func (recv *String) Equal(v2 *String) bool {
	c_v2 := (*C.GString)(C.NULL)
	if v2 != nil {
		c_v2 = (*C.GString)(v2.ToC())
	}

	retC := C.g_string_equal((*C.GString)(recv.native), c_v2)
	retGo := retC == C.TRUE

	return retGo
}

// Removes @len bytes from a #GString, starting at position @pos.
// The rest of the #GString is shifted down to fill the gap.
/*

C function

g_string_erase
*/
func (recv *String) Erase(pos int64, len int64) *String {
	c_pos := (C.gssize)(pos)

	c_len := (C.gssize)(len)

	retC := C.g_string_erase((*C.GString)(recv.native), c_pos, c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Frees the memory allocated for the #GString.
// If @free_segment is %TRUE it also frees the character data.  If
// it's %FALSE, the caller gains ownership of the buffer and must
// free it after use with g_free().
/*

C function

g_string_free
*/
func (recv *String) Free(freeSegment bool) string {
	c_free_segment :=
		boolToGboolean(freeSegment)

	retC := C.g_string_free((*C.GString)(recv.native), c_free_segment)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Creates a hash code for @str; for use with #GHashTable.
/*

C function

g_string_hash
*/
func (recv *String) Hash() uint32 {
	retC := C.g_string_hash((*C.GString)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Inserts a copy of a string into a #GString,
// expanding it if necessary.
/*

C function

g_string_insert
*/
func (recv *String) Insert(pos int64, val string) *String {
	c_pos := (C.gssize)(pos)

	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	retC := C.g_string_insert((*C.GString)(recv.native), c_pos, c_val)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Inserts a byte into a #GString, expanding it if necessary.
/*

C function

g_string_insert_c
*/
func (recv *String) InsertC(pos int64, c rune) *String {
	c_pos := (C.gssize)(pos)

	c_c := (C.gchar)(c)

	retC := C.g_string_insert_c((*C.GString)(recv.native), c_pos, c_c)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Inserts @len bytes of @val into @string at @pos.
// Because @len is provided, @val may contain embedded
// nuls and need not be nul-terminated. If @pos is -1,
// bytes are inserted at the end of the string.
//
// Since this function does not stop at nul bytes, it is
// the caller's responsibility to ensure that @val has at
// least @len addressable bytes.
/*

C function

g_string_insert_len
*/
func (recv *String) InsertLen(pos int64, val string, len int64) *String {
	c_pos := (C.gssize)(pos)

	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	c_len := (C.gssize)(len)

	retC := C.g_string_insert_len((*C.GString)(recv.native), c_pos, c_val, c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Converts a Unicode character into UTF-8, and insert it
// into the string at the given position.
/*

C function

g_string_insert_unichar
*/
func (recv *String) InsertUnichar(pos int64, wc rune) *String {
	c_pos := (C.gssize)(pos)

	c_wc := (C.gunichar)(wc)

	retC := C.g_string_insert_unichar((*C.GString)(recv.native), c_pos, c_wc)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds a string on to the start of a #GString,
// expanding it if necessary.
/*

C function

g_string_prepend
*/
func (recv *String) Prepend(val string) *String {
	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	retC := C.g_string_prepend((*C.GString)(recv.native), c_val)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Adds a byte onto the start of a #GString,
// expanding it if necessary.
/*

C function

g_string_prepend_c
*/
func (recv *String) PrependC(c rune) *String {
	c_c := (C.gchar)(c)

	retC := C.g_string_prepend_c((*C.GString)(recv.native), c_c)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Prepends @len bytes of @val to @string.
// Because @len is provided, @val may contain
// embedded nuls and need not be nul-terminated.
//
// Since this function does not stop at nul bytes,
// it is the caller's responsibility to ensure that
// @val has at least @len addressable bytes.
/*

C function

g_string_prepend_len
*/
func (recv *String) PrependLen(val string, len int64) *String {
	c_val := C.CString(val)
	defer C.free(unsafe.Pointer(c_val))

	c_len := (C.gssize)(len)

	retC := C.g_string_prepend_len((*C.GString)(recv.native), c_val, c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Converts a Unicode character into UTF-8, and prepends it
// to the string.
/*

C function

g_string_prepend_unichar
*/
func (recv *String) PrependUnichar(wc rune) *String {
	c_wc := (C.gunichar)(wc)

	retC := C.g_string_prepend_unichar((*C.GString)(recv.native), c_wc)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_string_printf : unsupported parameter ... : varargs

// Sets the length of a #GString. If the length is less than
// the current length, the string will be truncated. If the
// length is greater than the current length, the contents
// of the newly added area are undefined. (However, as
// always, string->str[string->len] will be a nul byte.)
/*

C function

g_string_set_size
*/
func (recv *String) SetSize(len uint64) *String {
	c_len := (C.gsize)(len)

	retC := C.g_string_set_size((*C.GString)(recv.native), c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Cuts off the end of the GString, leaving the first @len bytes.
/*

C function

g_string_truncate
*/
func (recv *String) Truncate(len uint64) *String {
	c_len := (C.gsize)(len)

	retC := C.g_string_truncate((*C.GString)(recv.native), c_len)
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Converts a #GString to uppercase.
/*

C function

g_string_up
*/
func (recv *String) Up() *String {
	retC := C.g_string_up((*C.GString)(recv.native))
	retGo := StringNewFromC(unsafe.Pointer(retC))

	return retGo
}

// An opaque data structure representing String Chunks.
// It should only be accessed by using the following functions.
/*

C type

GStringChunk
*/
type StringChunk struct {
	native *C.GStringChunk
}

func StringChunkNewFromC(u unsafe.Pointer) *StringChunk {
	c := (*C.GStringChunk)(u)
	if c == nil {
		return nil
	}

	g := &StringChunk{native: c}

	return g
}

func (recv *StringChunk) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Frees all memory allocated by the #GStringChunk.
// After calling g_string_chunk_free() it is not safe to
// access any of the strings which were contained within it.
/*

C function

g_string_chunk_free
*/
func (recv *StringChunk) Free() {
	C.g_string_chunk_free((*C.GStringChunk)(recv.native))

	return
}

// Adds a copy of @string to the #GStringChunk.
// It returns a pointer to the new copy of the string
// in the #GStringChunk. The characters in the string
// can be changed, if necessary, though you should not
// change anything after the end of the string.
//
// Unlike g_string_chunk_insert_const(), this function
// does not check for duplicates. Also strings added
// with g_string_chunk_insert() will not be searched
// by g_string_chunk_insert_const() when looking for
// duplicates.
/*

C function

g_string_chunk_insert
*/
func (recv *StringChunk) Insert(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_string_chunk_insert((*C.GStringChunk)(recv.native), c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Adds a copy of @string to the #GStringChunk, unless the same
// string has already been added to the #GStringChunk with
// g_string_chunk_insert_const().
//
// This function is useful if you need to copy a large number
// of strings but do not want to waste space storing duplicates.
// But you must remember that there may be several pointers to
// the same string, and so any changes made to the strings
// should be done very carefully.
//
// Note that g_string_chunk_insert_const() will not return a
// pointer to a string added with g_string_chunk_insert(), even
// if they do match.
/*

C function

g_string_chunk_insert_const
*/
func (recv *StringChunk) InsertConst(string string) string {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_string_chunk_insert_const((*C.GStringChunk)(recv.native), c_string)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// An opaque structure representing a test case.
/*

C type

GTestCase
*/
type TestCase struct {
	native *C.GTestCase
}

func TestCaseNewFromC(u unsafe.Pointer) *TestCase {
	c := (*C.GTestCase)(u)
	if c == nil {
		return nil
	}

	g := &TestCase{native: c}

	return g
}

func (recv *TestCase) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

/*

C type

GTestConfig
*/
type TestConfig struct {
	native          *C.GTestConfig
	TestInitialized bool
	TestQuick       bool
	TestPerf        bool
	TestVerbose     bool
	TestQuiet       bool
	TestUndefined   bool
}

func TestConfigNewFromC(u unsafe.Pointer) *TestConfig {
	c := (*C.GTestConfig)(u)
	if c == nil {
		return nil
	}

	g := &TestConfig{
		TestInitialized: c.test_initialized == C.TRUE,
		TestPerf:        c.test_perf == C.TRUE,
		TestQuick:       c.test_quick == C.TRUE,
		TestQuiet:       c.test_quiet == C.TRUE,
		TestUndefined:   c.test_undefined == C.TRUE,
		TestVerbose:     c.test_verbose == C.TRUE,
		native:          c,
	}

	return g
}

func (recv *TestConfig) ToC() unsafe.Pointer {
	recv.native.test_initialized =
		boolToGboolean(recv.TestInitialized)
	recv.native.test_quick =
		boolToGboolean(recv.TestQuick)
	recv.native.test_perf =
		boolToGboolean(recv.TestPerf)
	recv.native.test_verbose =
		boolToGboolean(recv.TestVerbose)
	recv.native.test_quiet =
		boolToGboolean(recv.TestQuiet)
	recv.native.test_undefined =
		boolToGboolean(recv.TestUndefined)

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : GTestLogBuffer

// Blacklisted : GTestLogMsg

// An opaque structure representing a test suite.
/*

C type

GTestSuite
*/
type TestSuite struct {
	native *C.GTestSuite
}

func TestSuiteNewFromC(u unsafe.Pointer) *TestSuite {
	c := (*C.GTestSuite)(u)
	if c == nil {
		return nil
	}

	g := &TestSuite{native: c}

	return g
}

func (recv *TestSuite) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// The #GThread struct represents a running thread. This struct
// is returned by g_thread_new() or g_thread_try_new(). You can
// obtain the #GThread struct representing the current thread by
// calling g_thread_self().
//
// GThread is refcounted, see g_thread_ref() and g_thread_unref().
// The thread represented by it holds a reference while it is running,
// and g_thread_join() consumes the reference that it is given, so
// it is normally not necessary to manage GThread references
// explicitly.
//
// The structure is opaque -- none of its fields may be directly
// accessed.
/*

C type

GThread
*/
type Thread struct {
	native *C.GThread
}

func ThreadNewFromC(u unsafe.Pointer) *Thread {
	c := (*C.GThread)(u)
	if c == nil {
		return nil
	}

	g := &Thread{native: c}

	return g
}

func (recv *Thread) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Waits until @thread finishes, i.e. the function @func, as
// given to g_thread_new(), returns or g_thread_exit() is called.
// If @thread has already terminated, then g_thread_join()
// returns immediately.
//
// Any thread can wait for any other thread by calling g_thread_join(),
// not just its 'creator'. Calling g_thread_join() from multiple threads
// for the same @thread leads to undefined behaviour.
//
// The value returned by @func or given to g_thread_exit() is
// returned by this function.
//
// g_thread_join() consumes the reference to the passed-in @thread.
// This will usually cause the #GThread struct and associated resources
// to be freed. Use g_thread_ref() to obtain an extra reference if you
// want to keep the GThread alive beyond the g_thread_join() call.
/*

C function

g_thread_join
*/
func (recv *Thread) Join() uintptr {
	retC := C.g_thread_join((*C.GThread)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// The #GThreadPool struct represents a thread pool. It has three
// public read-only members, but the underlying struct is bigger,
// so you must not copy this struct.
/*

C type

GThreadPool
*/
type ThreadPool struct {
	native *C.GThreadPool
	// _func : no type generator for Func, GFunc
	UserData  uintptr
	Exclusive bool
}

func ThreadPoolNewFromC(u unsafe.Pointer) *ThreadPool {
	c := (*C.GThreadPool)(u)
	if c == nil {
		return nil
	}

	g := &ThreadPool{
		Exclusive: c.exclusive == C.TRUE,
		UserData:  (uintptr)(c.user_data),
		native:    c,
	}

	return g
}

func (recv *ThreadPool) ToC() unsafe.Pointer {
	recv.native.user_data =
		(C.gpointer)(recv.UserData)
	recv.native.exclusive =
		boolToGboolean(recv.Exclusive)

	return (unsafe.Pointer)(recv.native)
}

// Frees all resources allocated for @pool.
//
// If @immediate is %TRUE, no new task is processed for @pool.
// Otherwise @pool is not freed before the last task is processed.
// Note however, that no thread of this pool is interrupted while
// processing a task. Instead at least all still running threads
// can finish their tasks before the @pool is freed.
//
// If @wait_ is %TRUE, the functions does not return before all
// tasks to be processed (dependent on @immediate, whether all
// or only the currently running) are ready.
// Otherwise the function returns immediately.
//
// After calling this function @pool must not be used anymore.
/*

C function

g_thread_pool_free
*/
func (recv *ThreadPool) Free(immediate bool, wait bool) {
	c_immediate :=
		boolToGboolean(immediate)

	c_wait_ :=
		boolToGboolean(wait)

	C.g_thread_pool_free((*C.GThreadPool)(recv.native), c_immediate, c_wait_)

	return
}

// Returns the maximal number of threads for @pool.
/*

C function

g_thread_pool_get_max_threads
*/
func (recv *ThreadPool) GetMaxThreads() int32 {
	retC := C.g_thread_pool_get_max_threads((*C.GThreadPool)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Returns the number of threads currently running in @pool.
/*

C function

g_thread_pool_get_num_threads
*/
func (recv *ThreadPool) GetNumThreads() uint32 {
	retC := C.g_thread_pool_get_num_threads((*C.GThreadPool)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Inserts @data into the list of tasks to be executed by @pool.
//
// When the number of currently running threads is lower than the
// maximal allowed number of threads, a new thread is started (or
// reused) with the properties given to g_thread_pool_new().
// Otherwise, @data stays in the queue until a thread in this pool
// finishes its previous task and processes @data.
//
// @error can be %NULL to ignore errors, or non-%NULL to report
// errors. An error can only occur when a new thread couldn't be
// created. In that case @data is simply appended to the queue of
// work to do.
//
// Before version 2.32, this function did not return a success status.
/*

C function

g_thread_pool_push
*/
func (recv *ThreadPool) Push(data uintptr) (bool, error) {
	c_data := (C.gpointer)(data)

	var cThrowableError *C.GError

	retC := C.g_thread_pool_push((*C.GThreadPool)(recv.native), c_data, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Sets the maximal allowed number of threads for @pool.
// A value of -1 means that the maximal number of threads
// is unlimited. If @pool is an exclusive thread pool, setting
// the maximal number of threads to -1 is not allowed.
//
// Setting @max_threads to 0 means stopping all work for @pool.
// It is effectively frozen until @max_threads is set to a non-zero
// value again.
//
// A thread is never terminated while calling @func, as supplied by
// g_thread_pool_new(). Instead the maximal number of threads only
// has effect for the allocation of new threads in g_thread_pool_push().
// A new thread is allocated, whenever the number of currently
// running threads in @pool is smaller than the maximal number.
//
// @error can be %NULL to ignore errors, or non-%NULL to report
// errors. An error can only occur when a new thread couldn't be
// created.
//
// Before version 2.32, this function did not return a success status.
/*

C function

g_thread_pool_set_max_threads
*/
func (recv *ThreadPool) SetMaxThreads(maxThreads int32) (bool, error) {
	c_max_threads := (C.gint)(maxThreads)

	var cThrowableError *C.GError

	retC := C.g_thread_pool_set_max_threads((*C.GThreadPool)(recv.native), c_max_threads, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Returns the number of tasks still unprocessed in @pool.
/*

C function

g_thread_pool_unprocessed
*/
func (recv *ThreadPool) Unprocessed() uint32 {
	retC := C.g_thread_pool_unprocessed((*C.GThreadPool)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Represents a precise time, with seconds and microseconds.
// Similar to the struct timeval returned by the gettimeofday()
// UNIX system call.
//
// GLib is attempting to unify around the use of 64bit integers to
// represent microsecond-precision time. As such, this type will be
// removed from a future version of GLib.
/*

C type

GTimeVal
*/
type TimeVal struct {
	native *C.GTimeVal
	TvSec  int64
	TvUsec int64
}

func TimeValNewFromC(u unsafe.Pointer) *TimeVal {
	c := (*C.GTimeVal)(u)
	if c == nil {
		return nil
	}

	g := &TimeVal{
		TvSec:  (int64)(c.tv_sec),
		TvUsec: (int64)(c.tv_usec),
		native: c,
	}

	return g
}

func (recv *TimeVal) ToC() unsafe.Pointer {
	recv.native.tv_sec =
		(C.glong)(recv.TvSec)
	recv.native.tv_usec =
		(C.glong)(recv.TvUsec)

	return (unsafe.Pointer)(recv.native)
}

// Adds the given number of microseconds to @time_. @microseconds can
// also be negative to decrease the value of @time_.
/*

C function

g_time_val_add
*/
func (recv *TimeVal) Add(microseconds int64) {
	c_microseconds := (C.glong)(microseconds)

	C.g_time_val_add((*C.GTimeVal)(recv.native), c_microseconds)

	return
}

// Opaque datatype that records a start time.
/*

C type

GTimer
*/
type Timer struct {
	native *C.GTimer
}

func TimerNewFromC(u unsafe.Pointer) *Timer {
	c := (*C.GTimer)(u)
	if c == nil {
		return nil
	}

	g := &Timer{native: c}

	return g
}

func (recv *Timer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Destroys a timer, freeing associated resources.
/*

C function

g_timer_destroy
*/
func (recv *Timer) Destroy() {
	C.g_timer_destroy((*C.GTimer)(recv.native))

	return
}

// If @timer has been started but not stopped, obtains the time since
// the timer was started. If @timer has been stopped, obtains the
// elapsed time between the time it was started and the time it was
// stopped. The return value is the number of seconds elapsed,
// including any fractional part. The @microseconds out parameter is
// essentially useless.
/*

C function

g_timer_elapsed
*/
func (recv *Timer) Elapsed(microseconds uint64) float64 {
	c_microseconds := (C.gulong)(microseconds)

	retC := C.g_timer_elapsed((*C.GTimer)(recv.native), &c_microseconds)
	retGo := (float64)(retC)

	return retGo
}

// This function is useless; it's fine to call g_timer_start() on an
// already-started timer to reset the start time, so g_timer_reset()
// serves no purpose.
/*

C function

g_timer_reset
*/
func (recv *Timer) Reset() {
	C.g_timer_reset((*C.GTimer)(recv.native))

	return
}

// Marks a start time, so that future calls to g_timer_elapsed() will
// report the time since g_timer_start() was called. g_timer_new()
// automatically marks the start time, so no need to call
// g_timer_start() immediately after creating the timer.
/*

C function

g_timer_start
*/
func (recv *Timer) Start() {
	C.g_timer_start((*C.GTimer)(recv.native))

	return
}

// Marks an end time, so calls to g_timer_elapsed() will return the
// difference between this end time and the start time.
/*

C function

g_timer_stop
*/
func (recv *Timer) Stop() {
	C.g_timer_stop((*C.GTimer)(recv.native))

	return
}

// Each piece of memory that is pushed onto the stack
// is cast to a GTrashStack*.
/*

C type

GTrashStack
*/
type TrashStack struct {
	native *C.GTrashStack
	// next : record
}

func TrashStackNewFromC(u unsafe.Pointer) *TrashStack {
	c := (*C.GTrashStack)(u)
	if c == nil {
		return nil
	}

	g := &TrashStack{native: c}

	return g
}

func (recv *TrashStack) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// The GTree struct is an opaque data structure representing a
// [balanced binary tree][glib-Balanced-Binary-Trees]. It should be
// accessed only by using the following functions.
/*

C type

GTree
*/
type Tree struct {
	native *C.GTree
}

func TreeNewFromC(u unsafe.Pointer) *Tree {
	c := (*C.GTree)(u)
	if c == nil {
		return nil
	}

	g := &Tree{native: c}

	return g
}

func (recv *Tree) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Removes all keys and values from the #GTree and decreases its
// reference count by one. If keys and/or values are dynamically
// allocated, you should either free them first or create the #GTree
// using g_tree_new_full(). In the latter case the destroy functions
// you supplied will be called on all keys and values before destroying
// the #GTree.
/*

C function

g_tree_destroy
*/
func (recv *Tree) Destroy() {
	C.g_tree_destroy((*C.GTree)(recv.native))

	return
}

// Unsupported : g_tree_foreach : unsupported parameter func : no type generator for TraverseFunc (GTraverseFunc) for param func

// Gets the height of a #GTree.
//
// If the #GTree contains no nodes, the height is 0.
// If the #GTree contains only one root node the height is 1.
// If the root node has children the height is 2, etc.
/*

C function

g_tree_height
*/
func (recv *Tree) Height() int32 {
	retC := C.g_tree_height((*C.GTree)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Inserts a key/value pair into a #GTree.
//
// If the given key already exists in the #GTree its corresponding value
// is set to the new value. If you supplied a @value_destroy_func when
// creating the #GTree, the old value is freed using that function. If
// you supplied a @key_destroy_func when creating the #GTree, the passed
// key is freed using that function.
//
// The tree is automatically 'balanced' as new key/value pairs are added,
// so that the distance from the root to every leaf is as small as possible.
/*

C function

g_tree_insert
*/
func (recv *Tree) Insert(key uintptr, value uintptr) {
	c_key := (C.gpointer)(key)

	c_value := (C.gpointer)(value)

	C.g_tree_insert((*C.GTree)(recv.native), c_key, c_value)

	return
}

// Gets the value corresponding to the given key. Since a #GTree is
// automatically balanced as key/value pairs are added, key lookup
// is O(log n) (where n is the number of key/value pairs in the tree).
/*

C function

g_tree_lookup
*/
func (recv *Tree) Lookup(key uintptr) uintptr {
	c_key := (C.gconstpointer)(key)

	retC := C.g_tree_lookup((*C.GTree)(recv.native), c_key)
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// Looks up a key in the #GTree, returning the original key and the
// associated value. This is useful if you need to free the memory
// allocated for the original key, for example before calling
// g_tree_remove().
/*

C function

g_tree_lookup_extended
*/
func (recv *Tree) LookupExtended(lookupKey uintptr, origKey uintptr, value uintptr) bool {
	c_lookup_key := (C.gconstpointer)(lookupKey)

	c_orig_key := (C.gpointer)(origKey)

	c_value := (C.gpointer)(value)

	retC := C.g_tree_lookup_extended((*C.GTree)(recv.native), c_lookup_key, &c_orig_key, &c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Gets the number of nodes in a #GTree.
/*

C function

g_tree_nnodes
*/
func (recv *Tree) Nnodes() int32 {
	retC := C.g_tree_nnodes((*C.GTree)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Removes a key/value pair from a #GTree.
//
// If the #GTree was created using g_tree_new_full(), the key and value
// are freed using the supplied destroy functions, otherwise you have to
// make sure that any dynamically allocated values are freed yourself.
// If the key does not exist in the #GTree, the function does nothing.
/*

C function

g_tree_remove
*/
func (recv *Tree) Remove(key uintptr) bool {
	c_key := (C.gconstpointer)(key)

	retC := C.g_tree_remove((*C.GTree)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// Inserts a new key and value into a #GTree similar to g_tree_insert().
// The difference is that if the key already exists in the #GTree, it gets
// replaced by the new key. If you supplied a @value_destroy_func when
// creating the #GTree, the old value is freed using that function. If you
// supplied a @key_destroy_func when creating the #GTree, the old key is
// freed using that function.
//
// The tree is automatically 'balanced' as new key/value pairs are added,
// so that the distance from the root to every leaf is as small as possible.
/*

C function

g_tree_replace
*/
func (recv *Tree) Replace(key uintptr, value uintptr) {
	c_key := (C.gpointer)(key)

	c_value := (C.gpointer)(value)

	C.g_tree_replace((*C.GTree)(recv.native), c_key, c_value)

	return
}

// Unsupported : g_tree_search : unsupported parameter search_func : no type generator for CompareFunc (GCompareFunc) for param search_func

// Removes a key and its associated value from a #GTree without calling
// the key and value destroy functions.
//
// If the key does not exist in the #GTree, the function does nothing.
/*

C function

g_tree_steal
*/
func (recv *Tree) Steal(key uintptr) bool {
	c_key := (C.gconstpointer)(key)

	retC := C.g_tree_steal((*C.GTree)(recv.native), c_key)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_tree_traverse : unsupported parameter traverse_func : no type generator for TraverseFunc (GTraverseFunc) for param traverse_func

// A utility type for constructing container-type #GVariant instances.
//
// This is an opaque structure and may only be accessed using the
// following functions.
//
// #GVariantBuilder is not threadsafe in any way.  Do not attempt to
// access it from more than one thread.
/*

C type

GVariantBuilder
*/
type VariantBuilder struct {
	native *C.GVariantBuilder
}

func VariantBuilderNewFromC(u unsafe.Pointer) *VariantBuilder {
	c := (*C.GVariantBuilder)(u)
	if c == nil {
		return nil
	}

	g := &VariantBuilder{native: c}

	return g
}

func (recv *VariantBuilder) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// #GVariantIter is an opaque data structure and can only be accessed
// using the following functions.
/*

C type

GVariantIter
*/
type VariantIter struct {
	native *C.GVariantIter
	// Private : x
}

func VariantIterNewFromC(u unsafe.Pointer) *VariantIter {
	c := (*C.GVariantIter)(u)
	if c == nil {
		return nil
	}

	g := &VariantIter{native: c}

	return g
}

func (recv *VariantIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Blacklisted : GVariantType
