<html>
<head>
</head>
<div id="header_title">
<h1>Profile</h1>
</div>

</nav>
</header>
<link rel="stylesheet" href="../assets/css/styleClear.css" type="text/css"> 
<body>
<div class="center">
</div>

<div class="box">
    <form action="/profile/compchangepasswd" method="POST">
        <input type="hidden" name="passwd" value="{{.Passwd}}">
        <input type="hidden" name="confirm" value="{{.Confirm}}">
        <input type="submit" value="change complete">
    </form>
</div>

</body>
</html>
