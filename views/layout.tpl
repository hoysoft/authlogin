<script src="/static/js/eldarion-ajax.js"></script>

{{if .LoginUser}}

    <div class="panel panel-primary" style="min-height:450px;height:auto!important; height:400px;">
        <div class="panel-heading"><i class="glyphicon glyphicon-wrench"></i>&nbsp;后台管理</div>
        <div class="panel-body">
       
                <div class="list-group " style="width:120px;float:left; height:100% "> 
                     <a href="/admin" class="list-group-item {{if compare  .__Tag  "admin" }}active {{end}}"><i class="glyphicon glyphicon-cog"></i>&nbsp;全局配置</a> 
                     <a href="/user" class="list-group-item {{if compare  .__Tag  "user" }}active {{end}}"><i class="glyphicon glyphicon-user"></i>&nbsp;{{.cnf.DefaultString "user::userManager" "User Manager"}}</a> 
                     <a href="#" class="list-group-item {{if compare  .__Tag  "" }}active {{end}}"><i class="glyphicon glyphicon-plus-sign"></i>&nbsp;角色管理</a> 
                     <a href="#" class="list-group-item {{if compare  .__Tag  "" }}active {{end}}"><i class="glyphicon glyphicon-ok-sign"></i>&nbsp;权限管理</a> 
					 <a href="/email" class="list-group-item {{if compare  .__Tag  "email" }}active {{end}}"><i class="glyphicon glyphicon-envelope"></i>&nbsp;Email配置</a>
                     <a href="/ldap" class="list-group-item {{if compare  .__Tag  "ldap" }}active {{end}}"><i class="glyphicon glyphicon-random"></i>&nbsp;{{.cnf.DefaultString "ldap::ldapManager" "LDAP Manager"}}</a>
					 <a href="#" class="list-group-item {{if compare  .__Tag  "" }}active {{end}}"><i class="glyphicon glyphicon-floppy-disk"></i>&nbsp;数据备份</a>
			    </div>
        
            <div class="col-md-10" >
		       {{.AdminContent}} 
		    </div>
        </div>
    </div>
{{else}}
{{.AdminContent}} 
{{end}}

 <div class="modal fade" id="popup_modal" tabindex="-1" role="dialog" aria-labelledby="popup_modal" aria-hidden="true">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
            <h4 class="modal-title" id="popup_modal_Title">###</h4>
          </div>
		  <div id="popup_modal_body">
		     <div id="popup_modal_content">
		 
		     </div>
		     <div class="modal-footer">
           
             </div>
		  </div>
        </div>
      </div>
</div>
 

<script language="javascript">
//通用模态框
$(".jexDlg").click(function(){
	//初始化为原始状态
	$("#popup_modal_body").html('<div id="popup_modal_content"></div><div class="modal-footer"> </div>');
	repeatId="#popup_modal_content"
	dataStr=this.getAttribute('data-dlg',2);
	if (dataStr){
	arr=dataStr.split(" ");
	//根据data-dlg替换不同方位区域
     repeatId= arr.indexOf('body')!==-1 ? "#popup_modal_body":"#popup_modal_content"
 
    }
	//	alert(repeatId);
 //	alert(this.getAttribute('href1',2));
    $("#popup_modal_Title").html(this.innerText );
//获取模态框显示数据
            $.ajax({url: this.getAttribute('href',2),
            type: 'GET',
            dataType: 'html', 
            timeout: 1000, 
 
            error: function(){
			 $("#popup_modal_content").html("error");
			},  
 
            success: function(result){
                $(repeatId).html(result);
            }  
 
            }); 
			
    //显示模态框 
     $('#popup_modal').modal({});
	
	return false;
	});

 
</script>