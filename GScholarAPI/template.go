package main

var homePagetmpl = `<html>
<head>
<title>Google Scholar API </title>
</head>
<body>
<h1> Google Scholar API </h1>
<!-- <p> This page allows you to upload a file (containing full names of users separated by newline) to retrieve google scholar meta data </p> -->
<p> Type a full name in the below text box to retrieve related google scholar meta data </p>
<form method="POST">
<label>Full Name: </label><br/>
<input type="text" name="fullName"><br />
<input type="submit">
</form>

<h3> Details are: </h3>  
    Full Name: &emsp; {{ .FullName }} <br/>
    Google Scholar Profile Link : &emsp; <a href="{{ .GScholarProfileLink }}"> {{ .FullName }} </a> <br/>
	Citations Count: &emsp; {{ .CitationsCount }} <br/>
	Designation: &emsp; {{ .Designation }} <br/>
	University/Affliation: &emsp; {{ .University }} <br/>
</body>
</html>

`