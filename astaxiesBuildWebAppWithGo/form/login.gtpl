<html>

<head>
	<title>Form</title>
</head>

<body>
	
	<form action="/login" method="post">
		Username: <input type="text" name="username" required /></br>
		Password: <input type="password" name="password" /></br>
		Email: <input type="email" name="email" /></br>
		Age: <input type="number" name="age" /></br>

		Gender: <select name="gender">
			<option value=""></option>
			<option value="male">Male</option>
			<option value="female">Female</option>
		</select><br/>
		<input type="hidden" name="token" value="{{.}}" />

		<input type="submit" value="Login" />
	</form>

</body>

</html>