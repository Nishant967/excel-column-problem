# Excel-Column-Problem
Finding column number from column title in excel.

###Note: 
1) By default http server wll run on port 9000.
2) API response will be text for now.
3) Converting input to uppercase string assuming there is no difference between uppercase and lowercase alphabet.
4) To run first do  go build solution.go and then run ./solution

#####Test API Endpoint : 
http://localhost:9000/v1/get-column-no?column-title=aa

Here column-title is the parameter where input need to be passed.

###ToDo List:
1) Writing unit test cases