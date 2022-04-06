{{define "base"}}
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    
    <link rel="stylesheet" type="text/css" href="/static/css/main.css" />
    {{block "style" .}} {{end}}
    <title>Document</title>
  </head>
  <body>
    {{block "content" .}} {{end}}
  </body>
</html>
{{end}}
