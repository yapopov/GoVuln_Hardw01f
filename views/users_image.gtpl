<html>
<head>
</head>
<div id="header_title">
<h1>User Image</h1>
</div>

</nav>
</header>
<link rel="stylesheet" href="../../assets/css/styleClear.css" type="text/css"> 
<body>
<div class="center">
</div>

<div class="box">
    <div class="showUpdateImage">
    {{ if .Image }}
    <img src="../../assets/img/{{.Image}}" width="500" height="400">
    {{ else }}
    <img src="../../assets/img/noimage.png" width="500" height="400">
    {{end}}
    </div>
    <form enctype="multipart/form-data" action="/profile/edit/upload" method="post">
  <input type="file" name="uploadfile" multiple="multiple" />
  <input type="submit" class="button_changeImage" value="Change Image" />
</div>

</body>
</html>
