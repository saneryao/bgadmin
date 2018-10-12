<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<meta charset="utf-8" />
<title>{{.PageTitle}}</title>
<meta name="description" content="overview &amp; stats" />
<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0" />
{{template "layout/header.tpl"}}
{{.CustomHeader}}
</head>
<body>
{{.CustomNavbar}}
<div class="main-container" id="main-container">
<script type="text/javascript">
	try{ace.settings.check('main-container' , 'fixed')}catch(e){}
</script>
{{.CustomSidebar}}
{{.LayoutContent}}
</div>
{{template "layout/footer.tpl"}}
{{.CustomFooter}}
</body>
</html>