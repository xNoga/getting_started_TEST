# getting_started_TEST

## Test cases

I have left out *Created By, Date of creation, Executed By and Date of execution* as I have created and executed all test on the same day; February 10, 2018.

Test Case ID | Test case description | Prerequisites | Test steps | Test data | Expected Result | Actual Result | Status | 
---|---|---|---|---|---|---|---|
TC001 | The objective of this test is to verify that a triangle with three sides of equal length will return 'equilateral' | The user should obtain three integers of equal length | Enter the three integers to the program and run it | The input should all be integers. Any other datatypes are not valid | 1. The program should return 'equilateral' if all three integers are equal to oneanother <br><br> 2. The program should return 'invalid input' if the input was not all integers | 1. The input was valid and the program returned 'equilateral' <br><br> 2. If the input was invalid the program will return 'invalid input' | Success |

Test Case ID | Test case description | Prerequisites | Test steps | Test data | Expected Result | Actual Result | Status | 
---|---|---|---|---|---|---|---|
TC002 | The objective of this test is to verify that a triangle with three sides of equal length will return 'equilateral' | The user should obtain three integers of equal length | Enter the three integers to the program and run it | The input should all be integers. Any other datatypes are not valid | 1. The program should return 'equilateral' if all three integers are equal to oneanother <br><br> 2. The program should return 'invalid input' if the input was not all integers | 1. The input was valid but the program **did not** return 'equilateral' <br><br> 2. If the input was invalid the program will return 'invalid input' | Fail |
