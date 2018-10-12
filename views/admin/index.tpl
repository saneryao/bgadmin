<title>{{i18n .Lang "site_name"}} {{i18n .Lang "back_short"}} - {{i18n .Lang "home"}}</title>

<div class="page-header">
	<div class="alert alert-block alert-success">
		<i class="ace-icon fa fa-check green"></i>
		{{i18n .Lang "welcome_to"}} <strong class="green">{{i18n .Lang "site_name"}} {{i18n .Lang "back_admin"}} {{i18n .Lang "system"}}<small>({{i18n .Lang "site_version"}})</small></strong> ! &nbsp;&nbsp; {{i18n .Lang "home_tip"}}
	</div>
</div><!-- /.page-header -->

<div class="row">
	<div class="col-xs-12">
		<div class="row">
			<div class="space-6"></div>

			<div class="col-sm-7 infobox-container">
				<!-- #section:pages/dashboard.infobox -->
				<div class="infobox infobox-green">
					<div class="infobox-icon">
						<i class="ace-icon fa fa-users"></i>
					</div>
					<div class="infobox-data">
						<span class="infobox-data-number" id="user_count">-</span>
						<div class="infobox-content">{{i18n .Lang "user"}}{{i18n .Lang "count"}}</div>
					</div>
				</div>

				<div class="infobox infobox-red">
					<div class="infobox-icon">
						<i class="ace-icon fa fa-user-secret"></i>
					</div>
					<div class="infobox-data">
						<span class="infobox-data-number" id="role_count">-</span>
						<div class="infobox-content">{{i18n .Lang "role"}}{{i18n .Lang "count"}}</div>
					</div>
				</div>

				<div class="infobox infobox-blue">
					<div class="infobox-icon">
						<i class="ace-icon fa fa-flask"></i>
					</div>
					<div class="infobox-data">
						<span class="infobox-data-number" id="resource_count">-</span>
						<div class="infobox-content">{{i18n .Lang "resource"}}{{i18n .Lang "count"}}</div>
					</div>
				</div>

				<div class="infobox infobox-red">
					
				</div>

				<div class="infobox infobox-orange2">
					<div class="infobox-chart">
						<span class="sparkline" data-values="176,135,170,128,191,153,202"></span>
					</div>
					<div class="infobox-data">
						<span class="infobox-data-number" id="article_count">-</span>
						<div class="infobox-content">{{i18n .Lang "article"}}{{i18n .Lang "count"}}</div>
					</div>
				</div>

				<div class="infobox infobox-blue2">
					<div class="infobox-chart">
						<span class="sparkline" data-values="196,128,202,94,100,170,224"></span>
					</div>
					<div class="infobox-data">
						<span class="infobox-data-number" id="comment_count">-</span>
						<div class="infobox-content">{{i18n .Lang "comment"}}{{i18n .Lang "count"}}</div>
					</div>
				</div>
				<div class="space-6"></div>

				<div class="infobox infobox-green infobox-small infobox-dark">
					<div class="infobox-progress">
						<div class="easy-pie-chart percentage" data-percent="61" data-size="39">
							<span class="percent" id="cpu_used">-</span>%
						</div>
					</div>
					<div class="infobox-data">
						<div class="infobox-content">{{i18n .Lang "cpu"}}</div>
						<div class="infobox-content">{{i18n .Lang "utilization"}}</div>
					</div>
				</div>

				<div class="infobox infobox-blue infobox-small infobox-dark">
					<div class="infobox-progress">
						<div class="easy-pie-chart percentage" data-percent="61" data-size="39">
							<span class="percent" id="memory_used">-</span>%
						</div>
					</div>
					<div class="infobox-data">
						<div class="infobox-content">{{i18n .Lang "memory"}}</div>
						<div class="infobox-content">{{i18n .Lang "utilization"}}</div>
					</div>
				</div>

				<div class="infobox infobox-grey infobox-small infobox-dark">
					<div class="infobox-progress">
						<div class="easy-pie-chart percentage" data-percent="61" data-size="39">
							<span class="percent" id="disk_used">-</span>%
						</div>
					</div>
					<div class="infobox-data">
						<div class="infobox-content">{{i18n .Lang "disk"}}</div>
						<div class="infobox-content">{{i18n .Lang "utilization"}}</div>
					</div>
				</div>
			</div>

			<div class="vspace-12-sm"></div>

			<div class="col-sm-5">
				<div class="widget-box">
					<div class="widget-header widget-header-flat widget-header-small">
						<h5 class="widget-title">
							{{i18n .Lang "article"}}{{i18n .Lang "distribution"}}
						</h5>
					</div>

					<div class="widget-body">
						<div class="widget-main">
							<div id="piechart-placeholder"></div>
							<div class="hr hr8 hr-double"></div>
							<div class="clearfix">
								<div class="grid2">
									<span class="grey">
										<i class="ace-icon fa fa-star fa-2x green"></i>
										&nbsp; {{i18n .Lang "favorite"}}{{i18n .Lang "count"}}
									</span>
									<h4 class="bigger pull-right"><span id="favorite_count">-</span>&nbsp;&nbsp;&nbsp;&nbsp;</h4>
								</div>

								<div class="grid2">
									<span class="grey">
										<i class="ace-icon fa fa-thumbs-o-up fa-2x blue"></i>
										&nbsp; {{i18n .Lang "praise"}}{{i18n .Lang "count"}}
									</span>
									<h4 class="bigger pull-right"><span id="praise_count">-</span>&nbsp;&nbsp;&nbsp;&nbsp;</h4>
								</div>
							</div>
						</div><!-- /.widget-main -->
					</div><!-- /.widget-body -->
				</div><!-- /.widget-box -->
			</div><!-- /.col -->
		</div><!-- /.row -->

		<div class="hr hr32 hr-dotted"></div>

		<div class="row">
			<div class="col-sm-12">
				<div class="widget-box transparent">
					<div class="widget-header widget-header-flat">
						<h4 class="widget-title lighter">
							<i class="ace-icon fa fa-eye"></i> {{i18n .Lang "visit"}}{{i18n .Lang "quantity"}}
						</h4>

						<div class="widget-toolbar">
							<a href="#" data-action="collapse">
								<i class="ace-icon fa fa-chevron-up"></i>
							</a>
						</div>
					</div>

					<div class="widget-body">
						<div class="widget-main padding-4">
							<div id="hits-charts"></div>
						</div><!-- /.widget-main -->
					</div><!-- /.widget-body -->
				</div><!-- /.widget-box -->
			</div><!-- /.col -->
		</div><!-- /.row -->

		<!-- PAGE CONTENT ENDS -->
	</div><!-- /.col -->
</div><!-- /.row -->

<!-- page specific plugin scripts -->

<!--[if lte IE 8]>
  <script src="/static/assets/js/excanvas.min.js"></script>
<![endif]-->
<script type="text/javascript">
	var scripts = [null,"/static/assets/js/jquery-ui.custom.min.js","/static/assets/js/jquery.ui.touch-punch.min.js","/static/assets/js/jquery.easypiechart.min.js","/static/assets/js/jquery.sparkline.index.min.js","/static/assets/js/jquery.flot.min.js","/static/assets/js/jquery.flot.pie.min.js","/static/assets/js/jquery.flot.resize.min.js","/static/assets/js/jquery.flot.time.js","/static/js/jquery.submit.js", null]
	$('.page-content-area').ace_ajax('loadScripts', scripts, function() {
		//inline scripts related to this page
		jQuery(function($) {
			$('.easy-pie-chart.percentage').each(function(){
				var $box = $(this).closest('.infobox');
				var barColor = $(this).data('color') || (!$box.hasClass('infobox-dark') ? $box.css('color') : 'rgba(255,255,255,0.95)');
				var trackColor = barColor == 'rgba(255,255,255,0.95)' ? 'rgba(255,255,255,0.25)' : '#E2E2E2';
				var size = parseInt($(this).data('size')) || 50;
				$(this).easyPieChart({
					barColor: barColor,
					trackColor: trackColor,
					scaleColor: false,
					lineCap: 'butt',
					lineWidth: parseInt(size/10),
					animate: ace.vars['old_ie'] ? false : 1000,
					size: size
				});
			})

			//
			$('.sparkline').each(function(){
				var $box = $(this).closest('.infobox');
				var barColor = !$box.hasClass('infobox-dark') ? $box.css('color') : '#FFF';
				$(this).sparkline('html', { tagValuesAttribute:'data-values',
														type: 'bar',
														barColor: barColor ,
														chartRangeMin:$(this).data('min') || 0
								 					});
			});

			//flot chart resize plugin, somehow manipulates default browser resize event to optimize it!
			//but sometimes it brings up errors with normal resize event handlers
			$.resize.throttleWindow = false;
			var placeholder = $('#piechart-placeholder').css({'width':'90%' , 'min-height':'150px'});
			function drawPieChart(placeholder, data, position) {
				$.plot(placeholder, data, {
					series: {
						pie: {
							show: true,
							tilt: 0.8,
							highlight: {
								opacity: 0.25
							},
							stroke: {
								color: '#fff',
								width: 2
							},
							startAngle: 2
						}
					},
					legend: {
						show: true,
						position: position || "ne",
						labelBoxBorderColor: null,
						margin: [-30,15]
					},
					grid: {
						hoverable: true,
						clickable: true
					}
				})
			}

			//pie chart tooltip example
			var $tooltip = $("<div class='tooltip top in'><div class='tooltip-inner'></div></div>").hide().appendTo('body');
			var previousPoint = null;
			placeholder.on('plothover', function (event, pos, item) {
				if (item) {
					if (previousPoint != item.seriesIndex) {
						previousPoint = item.seriesIndex;
						var tip = item.series['label'] + " : " + item.series['percent']+'%';
						$tooltip.show().children(0).text(tip);
					}
					$tooltip.css({top:pos.pageY + 10, left:pos.pageX + 10});
				} else {
					$tooltip.hide();
					previousPoint = null;
				}
			});

			/////////////////////////////////////
			$(document).one('ajaxloadstart.page', function(e) {
				$tooltip.remove();
			});

			var hits_charts = $('#hits-charts').css({'width':'100%' , 'height':'220px'});

			// 设置分布情况数据，进而描绘饼图
			function setDistributionsData(data) {
				drawPieChart(placeholder, data);
				placeholder.data('chart', data);
				placeholder.data('draw', drawPieChart);
			}

			// 设置访问量数据，进而描绘曲线图
			function setHitsData(data) {
				$.plot('#hits-charts', [
					{ label: '', data: data },
				], {
					// shadowSize: 0,
					series: {
						lines: { show: true },
						points: { show: true }
					},
					xaxis: {
						mode: 'time',
						timeformat: '%m-%d',
						tickSize: [2, 'day'],
						labelWidth: 45
					},
					// yaxis: {
					// },
					grid: {
						hoverable: true,
						backgroundColor: { colors: [ "#fff", "#fff" ] },
						borderWidth: 1,
						borderColor:'#555'
					}
				});
			}

			// 节点提示
			function showTooltip(x, y, contents) {
				$('<div id="tooltip">' + contents + '</div>').css( {
					position: 'absolute',
					display: 'none',
					top: y - 25,
					left: x - 8,
					border: '1px solid #fdd',
					padding: '2px',
					color: 'white',
					'background-color': 'black',
					opacity: 0.70
				}).appendTo("body").fadeIn(200);
			}

			// 绑定提示事件
			// var previousPoint = null;
			$("#hits-charts").bind("plothover", function (event, pos, item) {
				if (item) {
					if (previousPoint != item.dataIndex) {
						previousPoint = item.dataIndex;
						$("#tooltip").remove();
						var y = item.datapoint[1].toFixed(0);
						showTooltip(item.pageX, item.pageY,""+y);
					}
				} else {
					$("#tooltip").remove();
					previousPoint = null;
				}
			});

			////////////////////////Test//////////////////////////
			// var data1 = [
			// 	{ label: "social networks",  data: 38.7, color: "#68BC31"},
			// 	{ label: "search engines",  data: 24.5, color: "#2091CF"},
			// 	{ label: "ad campaigns",  data: 8.2, color: "#AF4E96"},
			// 	{ label: "direct traffic",  data: 18.6, color: "#DA5430"},
			// 	{ label: "other",  data: 10, color: "#FEE074"}
			// ];
			// setDistributionsData(data1);

			// var data2 = [];
			// var dateCurrent = new Date();
			// var tmCurrent = dateCurrent.getTime() + 28800000;
			// var tmToday = tmCurrent % 86400000;
			// tmCurrent -= tmToday;
			// for (var i = 0; i < 30; i ++) {
			// 	data2.push([(tmCurrent-86400000*i), Math.floor(Math.random()*100+1)]);
			// }
			// setHitsData(data2);

			submit_custom('{{urlfor "StatisticsApi.Get" ":entry" "amounts"}}', 'GET', [], function(data) {
				if (data.code) {
					$('#user_count').html(data.data.totalUsers);
					$('#role_count').html(data.data.totalRoles);
					$('#resource_count').html(data.data.totalResources);
					$('#article_count').html(data.data.totalArticles);
					$('#comment_count').html(data.data.totalComments);
				} else {
					alert(data.err);
				}
			});

			submit_custom('{{urlfor "StatisticsApi.Get" ":entry" "devices"}}', 'GET', [], function(data) {
				if (data.code) {
					$('#cpu_used').html(data.data.cpuUsed);
					$('#memory_used').html(data.data.memoryUsed);
					$('#disk_used').html(data.data.diskUsed);
				} else {
					alert(data.err);
				}
			});

			submit_custom('{{urlfor "StatisticsApi.Get" ":entry" "distributions"}}', 'GET', [], function(data) {
				if (data.code) {
					var dataDraw = [];
					var colorsDraw = ['#68BC31', '#2091CF', '#AF4E96', '#DA5430', '#FEE074', 'blue'];
					var i = 0;
					for (var key in data.data) {
						dataDraw.push( { label: key, data: data.data[key], color: colorsDraw[i++%6] } );
					}
					setDistributionsData(dataDraw);
					$('#favorite_count').html(data['totalFavorites']);
					$('#praise_count').html(data['totalPraises']);
				} else {
					alert(data.err);
				}
			});

			submit_custom('{{urlfor "StatisticsApi.Get" ":entry" "hits"}}', 'GET', [], function(data) {
				if (data.code) {
					var dataDraw = [];
					var dateCurrent = new Date();
					var tmCurrent = dateCurrent.getTime() + 28800000;
					var tmToday = tmCurrent % 86400000;
					tmCurrent -= tmToday;
					var i = 0;
					for (var key in data.data) {
						dataDraw.push( [ (tmCurrent-86400000*i), data.data[key] ]);
						i++;
					}
					setHitsData(dataDraw);
				} else {
					alert(data.err);
				}
			});


		})
	
	});
</script>
