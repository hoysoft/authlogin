<table class="user-list table table-striped table-hover">
<thead>
    <tr><th>用户名</th><th>姓氏</th><th>名字</th><th>E-mail</th><th>注册时间</th><th>最后登录</th><th>
	 <select class="form-horizontal form-control input-sm"  id="state" onchange="switchstate(this.value)" style="width: 78px;border-width:0px;background-color: transparent;margin:-2px -18px -2px -2px;" >
     {{range .userstates}}
        <option value={{.Id}} {{if .IsSelected}} select="selected"{{end}}>{{.Value}}</option>
     {{end}}
    </select>
	</th><th>管理</th></tr>
</thead>
<tbody>
    {{range .Users}}
    <tr><td>{{.Account}}</td><td>{{.LastName}}</td><td>{{.FirstName}}</td><td>{{.Email}}</td><td>{{.Createdtime.Format "2006-01-02 15:04:05" }}</td><td>{{.Lastlogintime.Format "2006-01-02 15:04:05"}}</td><td>{{.Status}}</td>
	<td><span class="action">
	<a  data-method="update"  href="/user/{{.Id}}">修改</a></span>&nbsp;<span class="action"><a href="/user/delete/{{.Id}}"   onclick="return confirm('确实要删除吗?')">删除</a></span>&nbsp;<span class="action"><a href="/user/delete/{{.Id}}">锁定</a></span></td></tr>
    {{end}}
</tbody>
</table><!-- End .user-list -->
{{template "authlogin/paginator.tpl" .}}