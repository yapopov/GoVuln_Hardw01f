<html>
<head>
</head>
<div id="header_title">
<h1>Change Image ?</h1>
</div>

</nav>
</header>
<link rel="stylesheet" href="./assets/css/styleClear.css" type="text/css"> 
<body>
<div class="center">
</div>

<div class="box">
    <div class="profileImage">
    {{ if .Image }}
    <img src="./assets/img/{{.Image}}" width="400" height="300">
    <h2>{{.Word}}</h2>
    {{ else }}
    <img src="./assets/img/noimage.png" width="400" height="300">
    {{end}}
    </div>
    <form enctype="text/plain" action="/profile/edit/upload" method="POST">
    <input type="file" name="uploadfile">
    <input type="submit" class="button_edit" value="upload">
    <div class="profileBox">
    </div>
</div>

</body>
</html>
