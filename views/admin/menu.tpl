<title>{{i18n .Lang "site_name"}} {{i18n .Lang "back_short"}} - {{i18n .Lang "menu"}}</title>

<div class="row">
	<div class="col-xs-12">
		<div class="table-header">
			<span class="white"><b>{{i18n .Lang "menu"}}{{i18n .Lang "list"}}</b></span>
		</div>

		<div>
			<table id="dynamic_table" class="table table-striped table-bordered table-hover row no-padding">
				<thead>
					<tr>
						<th>{{i18n .Lang "name"}}</th>
						<th>{{i18n .Lang "icon"}}</th>
						<th>{{i18n .Lang "url"}}</th>
						<th width="80">{{i18n .Lang "state"}}</th>
						<th width="80">{{i18n .Lang "operation"}}</th>
					</tr>
				</thead>

				<tbody class="row no-padding"></tbody>
			</table>
		</div>
	</div><!-- /.col -->
</div><!-- /.row -->

<a href="#modal-info" role="button" class="hidden" data-toggle="modal" id="btn-modal"></a>
<div id="modal-info" class="modal fade" tabindex="-1">
	<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header text-info bg-primary">
				<button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
				<h4 class="modal-title" id="info-title"></h4>
			</div>
			<div class="modal-body lead text-center no-padding" id="info-content">
				<form id="form-info">
				<fieldset>
					<div class="space-4"></div>
					<div class="row no-padding">
						<div class="col-sm-2 col-sm-offset-1">
							<label>{{i18n .Lang "name"}}<span class="red">*</span></label>
						</div>
						<div class="col-sm-8">
							<input type="text" name="name" class="form-control" maxlength="32" placeholder="{{i18n .Lang "menu"}}{{i18n .Lang "name"}}" required />
						</div>
					</div>

					<div class="space-4"></div>
					<div class="row">
						<div class="col-sm-2 col-sm-offset-1">
							<label>{{i18n .Lang "parent"}}{{i18n .Lang "menu"}}<span class="red">*</span></label>
						</div>
						<div class="col-sm-8">
							<select name="parent_id" class="form-control"></select>
						</div>
					</div>

					<div class="space-4"></div>
					<div class="row no-padding">
						<div class="col-sm-2 col-sm-offset-1">
							<label>{{i18n .Lang "icon"}}</label>
						</div>
						<div class="col-sm-8">
							<input type="text" name="icon" class="form-control" maxlength="32" placeholder="{{i18n .Lang "icon"}}" />
						</div>
					</div>

					<div class="space-4"></div>
					<div class="row">
						<div class="col-sm-2 col-sm-offset-1">
							<label>{{i18n .Lang "url"}}</label>
						</div>
						<div class="col-sm-8">
							<input type="text" name="url" class="form-control" maxlength="255" placeholder="{{i18n .Lang "url"}}" />
						</div>
					</div>

					<div class="space-4"></div>
					<div class="row">
						<div class="col-sm-2 col-sm-offset-1">
							<label>{{i18n .Lang "state"}}<span class="red">*</span></label>
						</div>
						<div class="col-sm-8">
							<select name="state" class="form-control" required>
								<option value="1">{{i18n .Lang "enable"}}</option>
								<option value="0">{{i18n .Lang "disable"}}</option>
							</select>
						</div>
					</div>

					<div class="space-4"></div>
					<div class="row">
						<div class="col-xs-7 col-xs-offset-4 col-sm-8 col-sm-offset-3 pull-left text-left">
							<button type="reset" class="btn btn-danger">
								<i class="ace-icon fa fa-refresh"></i>
								<span class="bigger-110">{{i18n .Lang "reset"}}</span>
							</button>&nbsp;&nbsp;&nbsp;&nbsp;
							<button type="sumbit" class="btn btn-success">
								<i class="ace-icon fa fa-check"></i>
								<span class="bigger-110">{{i18n .Lang "save"}}</span>
							</button>
						</div>
					</div>
					</fieldset>
				</form>
			</div>
		</div><!-- /.modal-content -->
	</div><!-- /.modal-dialog -->
</div>

<!-- page specific plugin scripts -->
<script type='text/javascript'>
	var scripts = [null,'/static/assets/js/jquery.dataTables.min.js','/static/assets/js/jquery.dataTables.bootstrap.min.js','/static/assets/js/dataTables.buttons.min.js','/static/assets/js/buttons.flash.min.js','/static/assets/js/buttons.html5.min.js','/static/assets/js/buttons.print.min.js','/static/assets/js/buttons.colVis.min.js','/static/js/jszip-2.5.0.min.js',/*'/static/js/pdfmake-0.1.36.min.js','/static/js/vfs_fonts--msyh.js',*/'/static/assets/js/dataTables.select.min.js','/static/assets/js/jquery-ui.custom.min.js','/static/assets/js/jquery.ui.touch-punch.min.js','/static/assets/js/chosen.jquery.min.js','/static/assets/js/spinbox.min.js','/static/assets/js/bootstrap-datepicker.min.js','/static/js/bootstrap-datepicker.locale.js','/static/js/jquery.validate.{{.Lang}}.js',null];
	var myTable;
	var myRow;
	function refresh_data() {
		if (myTable != null) {
			myTable.draw();
		}
	}
	function submit_callback(info) {
		if (info.code) {
			$('#modal-info').modal('hide');
			$('#msg-title').html('{{i18n .Lang "operate"}}{{i18n .Lang "success"}}');
			$('#msg-content').html(info.msg);
			$('#msg-dialog').modal('show');
			// alert(info.msg);
			if ($('#form-info').attr('method') == 'POST') {
				refresh_data();
			} else {
				if (myRow != null) {
					myRow.data(info.data);
				}
			}
		} else {
			// $('#msg-title').html('{{i18n .Lang "operate"}}{{i18n .Lang "failed"}}');
			// $('#msg-content').html(info.error);
			// $('#msg-dialog').modal('show');
			alert(info.error);
		}
		return false;
	}

	jQuery(function($) {
		$("#form-info").validate({
			errorElement:'div',
			errorPlacement: function(error, element) {
				if ($(element).next("div").hasClass("tooltip")) {
					$(element).attr("data-original-title", $(error).text()).tooltip("show");
				} else {
					$(element).attr("title", $(error).text()).tooltip("show");
				}
			},
			highlight: function(element, errorClass) {
				$(element).addClass(errorClass);
				$(element).parents('li:first').children('label').addClass(errorClass);
			},
			unhighlight:function(element, errorClass){
				$(element).removeClass(errorClass);
				$(element).parents('li:first').children('label').removeClass(errorClass);
			},
			submitHandler:function(form) {
				convpwd_submit(form, submit_callback);
				return false;
			}
		});
	});

	$('.page-content-area').ace_ajax('loadScripts', scripts, function() {
	 	//inline scripts related to this page
		jQuery(function($) {
			$('.date-picker').datepicker({
				language: '{{.Lang}}',
				autoclose: true,
				todayHighlight: true,
			})
			.next().on(ace.click_event, function(){
				$(this).prev().focus();
			});

			myTable =
			$('#dynamic_table').DataTable( {
				dom: '<"row"<"col-sm-8"B><"col-sm-4"f>>t<"row"<"col-sm-6"i><"col-sm-6"p>>',
				language: {
					url: '/static/js/dataTables.{{.Lang}}.json',
				},
				serverSide: true,
				ajax: {
					url: '{{urlfor "MenuAPI.Get"}}',
					type: 'GET',
					dataType: 'json',
					data: function(params) {
						return {
							draw: params.draw,
							limit: params.length,
							offset: params.start,
							per_page: params.length,
							search_column: 'name',
							search_value: params.search.value,
							sortby: 'id',
							order: 'asc',
						};
					},
					error: function (xhr, msg, e) {
						console.log(xhr.status);          // 状态码
						console.log(xhr.readyState);  // 状态
						console.log(msg);                 // 错误信息
						alert("" + xhr.status + msg);
					},
					complete: function (xhr, msg, e) {
						var obj = $('select[name="parent_id"]');
						obj.empty();
						obj.append('<option value="0">-</option>');
						var length = myTable.rows().data().length;
						for (var i = 0; i < length; i++) {
							var row = myTable.rows().data()[i];
							if (row.parent_id == 0) {
								obj.append('<option value="'+row.id+'">'+row.name+'</option>');
							}
						}
					},
				},
				processing: true,
				paging: true,
				lengthMenu: [1000],
				ordering: false,
				info: true,
				buttons: [
				  {
					text: '<span class="green"><i class="fa fa-plus bigger-110"></i> {{i18n .Lang "add"}}</span>',
					className: 'btn btn-white btn-primary btn-bold',
					columns: ':not(:first):not(:last)',
				  },
				  {
					text: '',
					className: 'btn btn-light disabled',
				  },
				  {
				  	text: '<i class="fa fa-refresh bigger-110 grey"></i> <span class="hidden">{{i18n .Lang "refresh"}}</span>',
				  	className: 'btn btn-white btn-primary btn-bold',
					action: function(e, dt, node, config) {
						refresh_data();
					},
				  },
				  {
					extend: 'colvis',
					text: '<i class="fa fa-viacoin bigger-110 blue"></i> <span class="hidden">{{i18n .Lang "show"}}/{{i18n .Lang "hide"}} {{i18n .Lang "columns"}}</span>',
					className: 'dt-button buttons-collection buttons-colvis btn btn-white btn-primary btn-bold',
					columns: ':not(:first):not(:last)',
				  },
				  {
					extend: 'pageLength',
					text: '<i class="fa fa-sort-amount-asc bigger-110 blue"></i> <span class="hidden">{{i18n .Lang "pagelen"}}</span>',
					className: 'dt-button buttons-collection buttons-copy btn btn-white btn-primary btn-bold',
				  },
				 //  {
					// extend: 'copy',
					// text: '<i class="fa fa-copy bigger-110 pink"></i> <span class="hidden">{{i18n .Lang "copy"}}{{i18n .Lang "to"}}{{i18n .Lang "clipboard"}}</span>',
					// className: 'dt-button buttons-collection buttons-copy btn btn-white btn-primary btn-bold',
				 //  },
				 //  {
					// extend: 'csv',
					// text: '<i class="fa fa-database bigger-110 orange"></i> <span class="hidden">{{i18n .Lang "export"}}{{i18n .Lang "to"}}{{i18n .Lang "csv"}}</span>',
					// className: 'dt-button buttons-collection buttons-csv btn btn-white btn-primary btn-bold',
					// exportOptions: {
					// 	columns: ':not(:last)',
					// },
				 //  },
				  {
					extend: 'excelHtml5',
					text: '<i class="fa fa-file-excel-o bigger-110 green"></i> <span class="hidden">{{i18n .Lang "export"}}{{i18n .Lang "to"}}{{i18n .Lang "excel"}}</span>',
					className: 'dt-button buttons-collection buttons-excel btn btn-white btn-primary btn-bold',
					exportOptions: {
						columns: ':not(:last)',
					},
				  },
				 //  {
					// extend: 'pdfHtml5',
					// text: '<i class="fa fa-file-pdf-o bigger-110 red"></i> <span class="hidden">{{i18n .Lang "export"}}{{i18n .Lang "to"}}{{i18n .Lang "pdf"}}</span>',
					// className: 'btn btn-white btn-primary btn-bold',
					// exportOptions: {
					// 	columns: ':not(:last)',
					// },
				 //  },
				  {
					extend: 'print',
					text: '<i class="fa fa-print bigger-110 grey"></i> <span class="hidden">{{i18n .Lang "print"}}</span>',
					className: 'dt-button buttons-collection buttons-print btn btn-white btn-primary btn-bold',
					autoPrint: false,
					// message: 'This print was produced using the Print button for DataTables',
					exportOptions: {
						columns: ':not(:last)',
					},
				  },
				],
				columns: [
					{ data: 'name', render: function(data, type, full) { if (full.parent_id==0) { return data; } else { return '&nbsp;&nbsp;&nbsp;&nbsp;└' +data; } } },
					{ data: 'icon', render: function(data) { if (data) { return '<i class="bigger-150 fa '+data+'"></i> &nbsp;' + data; } else { return ''; } } },
					{ data: 'url' },
					{ data: 'state', render: function(data) { if (data == 1) { return '{{i18n .Lang "enable"}}'; } return '{{i18n .Lang "disable"}}'; }  },
					{ data: -1 },
				],
				columnDefs: [ {
					targets: -1,
					data: 'Id',
					defaultContent: '<a id="edit" href="javascript:void(0);" class="green" data-rel="tooltip" title="{{i18n .Lang "edit"}}"><span class="blue"><i class="ace-icon fa fa-pencil-square-o bigger-120"></i></span></a>&nbsp;&nbsp;&nbsp;&nbsp;<a id="del" href="javascript:void(0);" data-rel="tooltip" title="{{i18n .Lang "del"}}"><span class="red"><i class="ace-icon fa fa-trash-o bigger-120"></i></span></a>',
				} ],
			} );

			// 数据编辑
			$('#dynamic_table tbody').on( 'click', 'a#edit', function () {
				myRow = myTable.row( $(this).parents('tr') );
				var data = myRow.data();
				$('#form-info').attr('action', '{{urlfor "MenuAPI.Put"}}/'+data.id);
				$('#form-info').attr('method', 'PUT');
				$('input[name="name"]').val(data.name);
				$('input[name="name"]').attr('value',data.name);
				$('select[name="parent_id"]').val(data.parent_id);
				$('select[name="parent_id"] option').removeAttr('selected');
				$('select[name="parent_id"] option[value="'+data.parent_id+'"]').attr('selected',true);
				$('input[name="icon"]').val(data.icon);
				$('input[name="icon"]').attr('value',data.icon);
				$('input[name="url"]').val(data.url);
				$('input[name="url"]').attr('value',data.url);
				$('select[name="state"]').val(data.state);
				$('select[name="state"] option').removeAttr('selected');
				$('select[name="state"] option[value="'+data.state+'"]').attr('selected',true);
				$('#info-title').html('{{i18n .Lang "edit"}}{{i18n .Lang "menu"}}{{i18n .Lang "page"}}');
				$('#form-info')[0].reset();
				$('#btn-modal').click();
			});

			// 数据删除
			$('#dynamic_table tbody').on( 'click', 'a#del', function () {
				if (confirm('确定删除吗？')) {
					var data = myTable.row( $(this).parents('tr') ).data();
					$.ajax({
						url: '{{urlfor "MenuAPI.Delete"}}/' + data['id'],
						type: 'DELETE',
						success: function (info) {
							if (info.code) {
								$('#msg-title').html('{{i18n .Lang "operate"}}{{i18n .Lang "success"}}');
								$('#msg-content').html(info.msg);
								$('#msg-dialog').modal('show');
								// alert(info.msg);
								refresh_data();
							} else {
								$('#msg-title').html('{{i18n .Lang "operate"}}{{i18n .Lang "failed"}}');
								$('#msg-content').html(info.error);
								$('#msg-dialog').modal('show');
								// alert(info.error);
							}
						},
						error: function (xhr, msg, e) {
							console.log(xhr.status);          // 状态码
							console.log(xhr.readyState);  // 状态
							console.log(msg);                 // 错误信息
							alert("" + xhr.status + msg);
						},
					})
				}
			});

			////
			setTimeout(function() {
				myTable.button(0).action(function (e, dt, button, config) {
					$('#form-info').attr('action', '{{urlfor "MenuAPI.Post"}}');
					$('#form-info').attr('method', 'POST');
					$('input[name="name"]').removeAttr('value');
					$('input[name="name"]').val('');
					$('select[name="parent_id"]').val(0);
					$('select[name="parent_id"] option').removeAttr('selected');
					$('select[name="parent_id"] option[value="0"]').attr('selected',true);
					$('input[name="icon"]').removeAttr('value');
					$('input[name="icon"]').val('');
					$('input[name="url"]').removeAttr('value');
					$('input[name="url"]').val('');
					$('select[name="state"]').val(1);
					$('select[name="state"] option').removeAttr('selected');
					$('select[name="state"] option[value="1"]').attr('selected',true);
					$('#info-title').html('{{i18n .Lang "add"}}{{i18n .Lang "menu"}}{{i18n .Lang "page"}}');
					$('#form-info')[0].reset();
					$('#btn-modal').click();
				});

				var defaultColvisAction = myTable.button(3).action();
				myTable.button(3).action(function (e, dt, button, config) {
					defaultColvisAction(e, dt, button, config);
					if($('.dt-button-collection > .dropdown-menu').length == 0) {
						$('.dt-button-collection')
						.wrapInner('<ul class="dropdown-menu dropdown-light dropdown-caret dropdown-caret" style="margin-left:150px;" />')
						.find('a').attr('href', '#').wrap("<li />")
					}
					$('.dt-button-collection').appendTo('.dt-buttons')
				});

				var defaultColvisAction = myTable.button(4).action();
				myTable.button(4).action(function (e, dt, button, config) {
					defaultColvisAction(e, dt, button, config);
					if($('.dt-button-collection > .dropdown-menu').length == 0) {
						$('.dt-button-collection')
						.wrapInner('<ul class="dropdown-menu dropdown-light dropdown-caret dropdown-caret" style="margin-left:200px;" />')
						.find('a').attr('href', '#').wrap("<li />")
					}
					$('.dt-button-collection').appendTo('.dt-buttons')
				});

				// var defaultCopyAction = myTable.button(5).action();
				// myTable.button(5).action(function (e, dt, button, config) {
				// 	defaultCopyAction(e, dt, button, config);
				// 	$('.dt-button-info').addClass('gritter-item-wrapper gritter-info gritter-center white');
				// });

				var noBtns = 0;
				$('.dt-buttons').find('a.dt-button').each(function() {
					if (noBtns > 1) {
						var div = $(this).find(' > div').first();
						if (div.length == 1) div.tooltip({container: 'body', title: div.parent().text()});
						else $(this).tooltip({container: 'body', title: $(this).text()});
					}
					noBtns++;
				});
			}, 600);
		});
	});
</script>