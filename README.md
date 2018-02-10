# getting_started_TEST

## Test cases

I have left out *Created By, Date of creation, Executed By and Date of execution* as I have created and executed all test on the same day; February 10, 2018.

Test Case ID | Test case description | Prerequisites | Test steps | Test data | Expected Result | Actual Result | Status | 
---|---|---|---|---|---|---|---|
TC001 | The objective of this test is to verify that a triangle with **three** sides of equal length will return 'equilateral triangle' | The user should obtain three equal integers | Enter the three integers to the program and run it | The input should all be integers. Any other datatypes are not valid | 1. The program should return 'equilateral triangle' if all three integers are equal to oneanother <br><br> 2. The program should return 'invalid input' if the input was not all integers | 1. The input was valid and the program returned 'equilateral triangle' <br><br> 2. If the input was invalid the program will return 'invalid input' | Success |

Test Case ID | Test case description | Prerequisites | Test steps | Test data | Expected Result | Actual Result | Status | 
---|---|---|---|---|---|---|---|
TC002 | The objective of this test is to verify that a triangle with **three** sides of equal length will return 'equilateral triangle' | The user should obtain three equal integers | Enter the three integers to the program and run it | The input should all be integers. Any other datatypes are not valid | 1. The program should return 'equilateral triangle' if all three integers are equal to oneanother <br><br> 2. The program should return 'invalid input' if the input was not all integers | 1. The input was valid but the program **did not** return 'equilateral triangle' <br><br> 2. If the input was invalid the program will return 'invalid input' | Fail |

Test Case ID | Test case description | Prerequisites | Test steps | Test data | Expected Result | Actual Result | Status | 
---|---|---|---|---|---|---|---|
TC003 | The objective of this test is to verify that a triangle with **no** sides of equal length will return 'scalene triangle' | The user should obtain three different integers | Enter the three integers to the program and run it | The input should all be integers. Any other datatypes are not valid | 1. The program should return 'scalene triangle' if all three integers are equal to oneanother <br><br> 2. The program should return 'invalid input' if the input was not all integers | 1. The input was valid and the program returned 'scalene triangle' <br><br> 2. If the input was invalid the program will return 'invalid input' | Success |

Test Case ID | Test case description | Prerequisites | Test steps | Test data | Expected Result | Actual Result | Status | 
---|---|---|---|---|---|---|---|
TC004 | The objective of this test is to verify that a triangle with **no** sides of equal length will return 'scalene triangle' | The user should obtain three different integers | Enter the three integers to the program and run it | The input should all be integers. Any other datatypes are not valid | 1. The program should return 'scalene triangle' if all three integers are equal to oneanother <br><br> 2. The program should return 'invalid input' if the input was not all integers | 1. The input was not valid and the program returned 'invalid input' <br><br> 2. If the input was invalid the program will return 'invalid input' | Fail |
