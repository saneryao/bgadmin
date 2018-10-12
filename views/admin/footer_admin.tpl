<!-- Custom Scripts -->
<script src="/static/assets/js/jquery.validate.min.js"></script>
<script src="/static/js/jquery.validate.{{.Lang}}.js"></script>
<!--
<script src="/static/js/jquery.serializejson.min.js"></script>
-->
<script src="/static/js/jquery.submit.js"></script>

<script src="/static/js/jsencrypt.min.js"></script>
<script src="/static/js/jquery.pwd.js"></script>

<!-- 底部copyright -->
<div class="footer">
  <div class="footer-inner">
    <div class="footer-content">
      <span class="bigger-120">
        <span class="blue bolder">{{i18n .Lang "site_name"}}</span>
        ({{i18n .Lang "site_domain"}}) &copy; 2018
      </span>

      &nbsp; &nbsp;
      <span class="action-buttons">
        <a href="#">
          <i class="ace-icon fa fa-twitter-square light-blue bigger-150"></i>
        </a>
        <a href="#">
          <i class="ace-icon fa fa-facebook-square text-primary bigger-150"></i>
        </a>
        <a href="#">
          <i class="ace-icon fa fa-rss-square orange bigger-150"></i>
        </a>
      </span>
    </div><!-- /.footer-content -->
  </div>
</div>

<script type="text/javascript">
	jQuery(function() {
		$('body').addClass("no-skin");
    // if(true) {
    //   $('#sidebar .nav-list > li').addClass('highlight');
    // } else {
    //   $('#sidebar .nav-list > li').removeClass('highlight');
    // }
    if(true) {
      $('#navbar').addClass('navbar-fixed-top');
    } else {
      $('#navbar').removeClass('navbar-fixed-top');
    }
    if(true) {
      $('#sidebar').addClass('sidebar-fixed');
      $('#menu-toggler').addClass('fixed');
    } else {
      $('#sidebar').removeClass('sidebar-fixed');
      $('#menu-toggler').removeClass('fixed');
    }
  });
</script>
