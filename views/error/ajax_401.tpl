<title>{{i18n .Lang "site_name"}} - 401</title>
<br/><br/>
<div class="bigger-160 center red">
  {{i18n .Lang "error_401"}}
</div>
<br/>
<div class="bigger-160 center red">
  <a href="{{urlfor "LogoutController.GET"}}">{{i18n .Lang "back_to"}}{{i18n .Lang "login"}}</a>
</div>