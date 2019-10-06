<html>
<head>
</head>
<div id="header_title">
<h1>Profile Edit</h1>
</div>

</nav>
</header>
<link rel="stylesheet" href="../assets/css/styleClear.css" type="text/css"> 
<body>
<div class="center">
</div>

<div class="box">
    <div class="profileImage">
    {{ if .Image }}
    <img src="../assets/img/{{.Image}}" width="400" height="300">
    {{ else }}
    <img src="../assets/img/noimage.png" width="400" height="300">
    {{end}}
    </div>
    <div class="profileBox">
    <form action="/profile/edit/confirm" method="post">
    <h1>Name : <t> <input type="text" name="username" value="{{.UserName}}"> </h1>
    <h2>Age : <t> <input type="text" name="age" value="{{.Age}}"> </h2>
    <h2>Mail : <t> <input type="text" name="mail" value="{{.Mail}}" disabled="disable"> </h2>
    <h2>Address : <t> <input type="text" name="address" value="{{.Address}}"></h2>
    <h2>Favorite Animal : <t> <input type="text" name="animal" value="{{.Animal}}"> </h2>
    <h2>Word : <t> <input type="text" name="word" value="{{.Word}}"> </h2>
    </div>
    <input type="submit" class="button_edit" value="Confirm">
</div>

</body>
</html>
