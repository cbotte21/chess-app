/*
* This is a VERY basic length validator. Will soon implement rigid string validation.
 */

package utilities

/*
*	function: ParseEmail
*	parameter: a string for a candide email
*	result: a boolean validating a email
 */
func ParseEmail(username string) bool {
	//Should include profanity detection
	return len(username) >= 4
}

/*
*	function: ParsePassword
*	parameter: a string for a candite password
*	result: a boolean validating a password
 */
func ParsePassword(password string) bool {
	return len(password) >= 7
}
