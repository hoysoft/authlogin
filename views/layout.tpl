{{if .LoginUser}}
    <div class="panel panel-primary" style="min-height:450px;height:auto!important; height:400px;">
        <div class="panel-heading"><i class="glyphicon glyphicon-wrench"></i>&nbsp;后台管理</div>
        <div class="panel-body">
       
                <div class="list-group col-md-2"> 
                     <a href="/admin" class="list-group-item {{if compare  .__Tag  "admin" }}active {{end}}"><i class="glyphicon glyphicon-cog"></i>&nbsp;全局配置</a> 
                     <a href="/user" class="list-group-item {{if compare  .__Tag  "user" }}active {{end}}"><i class="glyphicon glyphicon-user"></i>&nbsp;{{.cnf.DefaultString "user::userManager" "User Manager"}}</a> 
                     <a href="#" class="list-group-item {{if compare  .__Tag  "" }}active {{end}}"><i class="glyphicon glyphicon-plus-sign"></i>&nbsp;角色管理</a> 
                     <a href="#" class="list-group-item {{if compare  .__Tag  "" }}active {{end}}"><i class="glyphicon glyphicon-ok-sign"></i>&nbsp;权限管理</a> 
                     <a href="/ldap" class="list-group-item {{if compare  .__Tag  "ldap" }}active {{end}}"><i class="glyphicon glyphicon-random"></i>&nbsp;{{.cnf.DefaultString "ldap::ldapManager" "LDAP Manager"}}</a>
					 <a href="#" class="list-group-item {{if compare  .__Tag  "" }}active {{end}}"><i class="glyphicon glyphicon-floppy-disk"></i>&nbsp;数据备份</a>
			    </div>
        
            <div class="col-md-10">
		       {{.AdminContent}} 
		    </div>
        </div>
    </div>
{{else}}
{{.AdminContent}} 
{{end}}


 


