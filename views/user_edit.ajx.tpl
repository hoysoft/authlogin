
		<div class="col-md-12">			
			<form class="form-horizontal templatemo-create-account templatemo-container" role="form" action="#" method="post">
				<div class="form-inner">
					<div class="form-group">
					  <div class="col-md-6">		          	
			            <label for="last_name" class="control-label">{{.cnf.DefaultString "user::Lastname" "Last Name"}}</label>
			            <input type="text" class="form-control" id="last_name" placeholder="">		            		            		            
			          </div>   
			          <div class="col-md-6">		          	
			            <label for="first_name" class="control-label">{{.cnf.DefaultString "user::Firstname" "First Name"}}</label>
			            <input type="text" class="form-control" id="first_name" placeholder="">		            		            		            
			          </div>  
			                    
			        </div>
			        <div class="form-group">
			          <div class="col-md-12">		          	
			            <label for="username" class="control-label">Email</label>
			            <input type="email" class="form-control" id="email" placeholder="" onchange="checkStatus()">		            		            		            
			          </div>              
			        </div>			
			        <div class="form-group">
			          <div class="col-md-6">		          	
			            <label for="username" class="control-label">{{.cnf.DefaultString "user::Username" "Username"}}</label>
			            <input type="text" class="form-control" id="username" placeholder="" onchange="checkStatus()">		            		            		            
			          </div>
			          <div class="col-md-6 templatemo-radio-group">
			          	<label class="radio-inline">
		          			<input type="radio" name="optionsRadios" id="optionsRadios1" value="option1"> {{.cnf.DefaultString "user::Male" "Male"}}
		          		</label>
		          		<label class="radio-inline">
		          			<input type="radio" name="optionsRadios" id="optionsRadios2" value="option2"> {{.cnf.DefaultString "user::Female" "Female"}}
		          		</label>
			          </div>             
			        </div>
					
					{{if .IsCreateUser}}
			        <div class="form-group">
			          <div class="col-md-6">
			            <label for="password" class="control-label">{{.cnf.DefaultString "user::Password" "Password"}}</label>
			            <input type="password" class="form-control" id="password" placeholder="" onchange="checkStatus()">
			          </div>
			          <div class="col-md-6">
			            <label for="password" class="control-label">{{.cnf.DefaultString "user::ConfirmPassword" "Confirm Password"}}</label>
			            <input type="password" class="form-control" id="password_confirm" placeholder="" onchange="checkStatus()">
			          </div>
			        </div>
					
			        <div class="form-group">
			          <div class="col-md-12">
			            <label><input id="_checkbox" type="checkbox" onclick="checkStatus()">{{.cnf.DefaultString "user::IagreeTo" "I agree to the"}} 
						        <a href="/static/mm1.html" data-toggle="modal" data-target="#templatemo_modal">
								{{.cnf.DefaultString "user::TermsOfService" "Terms of Service"}}</a> {{.cnf.DefaultString "default::and" "and"}} 
								 <a href="#">{{.cnf.DefaultString "user::PrivacyPolicy" "Privacy Policy"}}.</a></label>
			          </div>
			        </div>
					
					{{end}}
			        <div class="form-group">
			          <div class="col-md-12">
			            <input type="submit" id="_btnsubmit" value='{{.cnf.DefaultString "user::CreateAccount" "Create account"}}' class="btn btn-info" disabled="true">
			         {{if .IsCreateUser}}   
					   <a href="/user/login" class="pull-right">{{.cnf.DefaultString "user::btnLogin" "Login"}}</a>
					{{end}}
			          </div>
			        </div>	
				</div>				    	
		      </form>		      
 
	</div>
	
	<!-- Modal -->
	<div class="modal fade" id="templatemo_modal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
	  <div class="modal-dialog">
	    <div class="modal-content">
	      
	    </div>
	  </div>
	</div>
	
	{{template "authlogin/templatemo_style.css.tpl" .}}
 
<script type="text/javascript"> 

//验证email地址
function IsMail(mail)
{ 
  var reMail = /^(?:[a-z\d]+[_\-\+\.]?)*[a-z\d]+@(?:([a-z\d]+\-?)*[a-z\d]+\.)+([a-z]{2,})+$/i;
  return reMail.test(mail);
 }
	
 function checkStatus() {	
   v=document.getElementById("_checkbox");
   if (v.checked && 
       document.getElementById("username").value!="" &&
        IsMail(document.getElementById("email").value) &&
       document.getElementById("password").value==document.getElementById("password_confirm").value)
    {
	  document.getElementById("_btnsubmit").removeAttribute("disabled");
    }else {
	  document.getElementById("_btnsubmit").disabled="true";
	};
 }

</script>