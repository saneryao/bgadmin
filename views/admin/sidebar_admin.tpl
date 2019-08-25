<div id="sidebar" class="sidebar responsive ace-save-state">
	<script type="text/javascript">
		try{ace.settings.loadState('sidebar')}catch(e){}
	</script>

	<div class="sidebar-shortcuts" id="sidebar-shortcuts">
		<div class="sidebar-shortcuts-large" id="sidebar-shortcuts-large">
			<!--
			<button class="btn btn-success">
				<i class="ace-icon fa fa-signal"></i>
			</button>
			<button class="btn btn-info">
				<i class="ace-icon fa fa-pencil"></i>
			</button>
			<button class="btn btn-warning">
				<i class="ace-icon fa fa-users"></i>
			</button>
			<button class="btn btn-danger">
				<i class="ace-icon fa fa-cogs"></i>
			</button>
			-->
			<img src="/static/img/logo.png" class="sidebar-collapse" style="border-radius:50%;border: 2px solid grey" width="100" height="100" />
		</div>

		<div class="sidebar-shortcuts-mini" id="sidebar-shortcuts-mini">
			<!--
			<span class="btn btn-success"></span>
			<span class="btn btn-info"></span>
			<span class="btn btn-warning"></span>
			<span class="btn btn-danger"></span>
			-->
			<img src="/static/img/logo.png" class="sidebar-collapse" style="border-radius:50%;border: 2px solid grey" width="42" height="42" />
		</div>
	</div><!-- /.sidebar-shortcuts -->

	<ul class="nav nav-list">
		<li class="">
			<a data-url="page/index" href="#page/index">
				<i class="menu-icon fa fa-home"></i>
				<span class="menu-text"> {{i18n .Lang "home"}} </span>
			</a>
			<b class="arrow"></b>
		</li>
		<b class="arrow"></b>

		<li class="">
			<a href="#" class="dropdown-toggle">
				<i class="menu-icon fa fa-users"></i>
				<span class="menu-text"> {{i18n .Lang "user_and_power"}} </span>
				<b class="arrow fa fa-angle-down"></b>
			</a>
			<b class="arrow"></b>

			<ul class="submenu">
				<li class="">
					<a data-url="page/user" href="#page/user">
						<i class="menu-icon fa fa-caret-right"></i>{{i18n .Lang "user"}}
					</a>
					<b class="arrow"></b>
				</li>
				<b class="arrow"></b>

				<li class="">
					<a data-url="page/role" href="#page/role">
						<i class="menu-icon fa fa-caret-right"></i>{{i18n .Lang "role"}}
					</a>
					<b class="arrow"></b>
				</li>
				<b class="arrow"></b>

				<li class="">
					<a data-url="page/menu" href="#page/menu">
						<i class="menu-icon fa fa-caret-right"></i>{{i18n .Lang "menu"}}
					</a>
					<b class="arrow"></b>
				</li>
				<b class="arrow"></b>

				<li class="">
					<a data-url="page/link" href="#page/link">
						<i class="menu-icon fa fa-caret-right"></i>{{i18n .Lang "api"}}
					</a>
					<b class="arrow"></b>
				</li>
				<b class="arrow"></b>
			</ul>
		</li>
		<b class="arrow"></b>

		<li class="">
			<a href="#" class="dropdown-toggle">
				<i class="menu-icon fa fa-info-circle"></i>
				<span class="menu-text"> {{i18n .Lang "more_demos"}} </span>
				<b class="arrow fa fa-angle-down"></b>
			</a>
			<b class="arrow"></b>

			<ul class="submenu">
				<li class="">
					<a data-url="page/demo" href="#page/demo">
						<i class="menu-icon fa fa-caret-right"></i>{{i18n .Lang "page"}}
					</a>
					<b class="arrow"></b>
				</li>
				<b class="arrow"></b>

				<li class="">
					<a data-url="page/grid" href="#page/grid">
						<i class="menu-icon fa fa-caret-right"></i>{{i18n .Lang "grid"}}
					</a>
					<b class="arrow"></b>
				</li>
				<b class="arrow"></b>

				<li class="">
					<a data-url="page/button" href="#page/button">
						<i class="menu-icon fa fa-caret-right"></i>{{i18n .Lang "button"}}
					</a>
					<b class="arrow"></b>
				</li>
				<b class="arrow"></b>

				<li class="">
					<a data-url="page/icon" href="#page/icon">
						<i class="menu-icon fa fa-caret-right"></i>{{i18n .Lang "icon"}}
					</a>
					<b class="arrow"></b>
				</li>
				<b class="arrow"></b>
			</ul>
		</li>
		<b class="arrow"></b>
	</ul><!-- /.nav-list -->

	<div class="sidebar-footer text-center green">
		<div class="hr hr32 hr-dotted"></div>
		<span class="bigger-110">{{i18n .Lang "site_version"}}</span>
	</div>
</div>