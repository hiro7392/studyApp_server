<!DOCTYPE html>
<html lang="ja">
<head>
  <meta charset="UTF-8">
  <title>{{ .Title }}</title>
</head>
<body>
  <h1>{{ .Title }}</h1>
  <form action="/sigin" method="POST">
    <input type="text" name="name">
    <button type="submit">送信</button>
  </form>
  <p>&copy;{{ .Footer }}</p>
</body>
</html>