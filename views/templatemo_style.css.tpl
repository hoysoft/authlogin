<style>
/* 


Credit: www.cssmoban.com
	1. General
	2. login-form-1
	3. login-form-2
	4. inline-login
	5. create-account
	6. forgot-password
	7. contact-form-1
	8. contact-form-2
	9. payment-form
	10. Media Queries
	
--------------------------------------------*/
/* 1. General
--------------------------------------------*/

*, body { font-family: "HelveticaNeue", "Helvetica Neue", Helvetica, Arial, sans-serif; }
body { margin-bottom: 30px; }
h1 {
	font-weight: 300;
	font-size: 36px;
	line-height: normal;
	margin: auto;
	margin-top: 30px;
	text-align: center;
}
h2 {
	font-weight: 500;
	font-size: 20px;
	line-height: normal;
	margin: auto;
	margin-bottom: 30px;
	text-align: center;
	color: #09F;
}

.green { color: #3C3; }
.gray { color: #999; }

.form-horizontal .control-label { padding-top: 0; }
.form-horizontal.templatemo-payment-form .radio.templatemo-no-padding-top {	padding-top: 0; }
.control-wrapper {
	position: relative;
	padding-left: 30px;
}
.control-wrapper label.fa-label {
	position: absolute;
	left: 4px;
	top: 6px;
}
.form-group:last-child { margin-bottom: 0; }
.form-horizontal .control-label { margin-bottom: 10px; }
.table { margin-bottom: 0; }
.table>thead>tr>th, .table>tbody>tr>th, .table>tfoot>tr>th, 
.table>thead>tr>td, .table>tbody>tr>td, .table>tfoot>tr>td {
	vertical-align: middle;
}
.templatemo-form-list-container { max-width: 600px; }
.templatemo-input-icon-container { position: relative; }
.templatemo-input-icon-container .fa {
	color: gray;
	position: absolute;
	left: 10px;
	top: 10px;
}
.templatemo-input-icon-container input, 
.templatemo-input-icon-container textarea {
	padding-left: 30px;
}
.templatemo-container {
	background-color: rgba(255,255,255,0.8);
	border: 1px solid #dedede;
	border-radius: 8px;
	margin: 0 auto;
	padding: 30px;
}
.templatemo-bg-gray { background-color: #eee; }
.templatemo-bg-gray h1 { color: rgb(74, 164, 180); }
.templatemo-bg-image-1 {
	background-color: rgb(60,60,60);
	background-image: url(../images/templatemo-bg-1.jpg);	
}
.templatemo-bg-image-2 {
	background-color: rgb(70, 90, 40);
	background-image: url(../images/templatemo-bg-2.jpg);
}
.templatemo-bg-image-1, .templatemo-bg-image-2 {
	background-repeat: no-repeat;
	background-position: center center;
	background-attachment: fixed;	
}
.templatemo-bg-image-1, .templatemo-bg-image-2 { background-size: cover; }
.font-size-small { font-size: 0.8em; }
.margin-bottom-15 {	margin-bottom: 15px; }
.margin-bottom-30 {	margin-bottom: 30px; }
.form-group { margin-bottom: 20px; }
.form-group a {	line-height: 34px; }
.fa { font-size: 16px; }
.fa.login-with {
	font-size: 30px;
	margin: 0 5px;
}
.fa-medium { font-size: 20px; }
.inline-block { display: inline-block; }


/*-------------------------------------
2. login-form-1
---------------------------------------*/
.templatemo-login-form-1 { max-width: 500px; }
.templatemo-login-form-1 a { color: gray; }
.templatemo-login-form-1 a:hover {
	color: black;
	text-decoration: none;
	cursor: pointer;
}
.templatemo-create-new {
	color: #58B4BB;
	font-size: 18px;
	font-weight: 300;
}
.templatemo-create-new:hover {
	color: #138892;
	text-decoration: none;
}

/*-------------------------------------
3. login-form-2
---------------------------------------*/
.templatemo-login-form-2 {
	background-color: rgba(13,13,13,0.25);
	border-radius: 8px;
	color: #fff;
	font-weight: 300;
	max-width: 650px;
	padding: 0 30px 30px 30px;
	margin: 30px auto 0 auto;
}
.templatemo-login-form-2 h1 { color: #fff; margin-bottom: 40px; }
.templatemo-login-form-2 a { color: #DBDBDB; }
.templatemo-login-form-2 .form-control {
	background-color: rgba(83, 78, 78, 0.35);
	border: 1px solid rgba(255, 255, 255, 0.27);
	color: rgb(255,255,255);
}
.templatemo-login-form-2 .templatemo-input-icon-container .fa {	color: rgb(190, 190, 190); }
.templatemo-login-form-2 .control-label { margin-bottom: 10px; }
.templatemo-login-form-2 .btn {	width: 100%; }
.templatemo-login-form-2 .btn-social { max-width: 220px; }
.templatemo-login-form-2 label { font-weight: 400; }
.templatemo-login-form-2 .templatemo-one-signin { border-right: 1px solid rgba(200, 200, 200, 0.5); }


/*-------------------------------------
4. inline-login
--------------------------------------*/
.templatemo-header {
	background-color: #0B8290;
	color: white;
	min-height: 90px;
}
.templatemo-header .checkbox {
	display: block;
	margin-top: 5px;
}
.templatemo-header form { margin-top: 25px; }
.templatemo-header .form-inline .form-group { vertical-align: top; }

.logo { margin: 18px; }


/*-------------------------------------
5. create-account
---------------------------------------*/
.templatemo-create-account {
	background-color: #fff;
	border: none;
/*	border: 1px solid rgba(176, 176, 176, 0.4);*/
	max-width: 700px;	
}
.templatemo-create-account label { font-weight: 400; }
.templatemo-create-account input[type='checkbox'] {
	margin-right: 10px;
	position: relative;
	top: -1.5px;
}
.templatemo-radio-group { margin-top: 30px; }


/*-------------------------------------
6. forgot-password
---------------------------------------*/
.templatemo-forgot-password-form {
	background-color: #fff;	
	max-width: 550px;
	margin: 0 auto;
	padding: 30px;
}


/*-------------------------------------
7. contact-form-1
---------------------------------------*/
.templatemo-contact-form-1 {
	background: rgba(0,0,0,0.6);
	border-radius: 8px;
	color: rgb(197,197,197);
	max-width: 600px;
	margin: 30px auto 0 auto;
	padding: 0 30px 30px 30px;
}
.templatemo-contact-form-1 h1, .templatemo-contact-form-1 p { color: rgb(255,255,255); }
.templatemo-contact-form-1 a {
	color: #FF9;
	text-decoration: underline;
}
.templatemo-contact-form-1 a:hover { color: rgb(255,255,255); }
.templatemo-contact-form-1 label { font-weight: 400;}
.templatemo-contact-form-1 .form-control {
	font-size: 16px;
	padding: 10px 12px 10px 35px;
}
.templatemo-contact-form-1 input.form-control {	height: 45px; }
.templatemo-contact-form-1 .fa { font-size: 20px; }
.templatemo-contact-form-1 .form-control {
	background-color: rgba(0, 0, 0, 0.6);
	border: 1px solid rgba(176, 176, 176, 0.4);
	color: white;
}
.templatemo-contact-form-1 .form-control:focus {
	box-shadow: inset 0 1px 1px rgba(140, 220, 60, 0.7),0 0 10px rgba(140, 220, 60, 0.7);
}
.templatemo-contact-form-1 .templatemo-input-icon-container .fa {
	color: rgb(197,197,197);
	left: 12px;
	top: 12px;
}
.templatemo-contact-form-1 .templatemo-input-icon-container .fa-envelope-o {
	font-size: 18px;
	top: 14px;
}


/*-------------------------------------
8. contact form 2
---------------------------------------*/
.templatemo-contact-form-2 {
	border-radius: 8px;
	max-width: 960px;
	margin: 0 auto;
	padding: 30px 30px 0 30px;
}
.templatemo-contact-form-2 .form-group { margin-bottom: 30px; }
.templatemo-contact-form-2 label { font-weight: 400; }
.templatemo-contact-form-2 textarea { min-height: 225px; }


/*-------------------------------------
9. payment-form
---------------------------------------*/
.templatemo-payment-form {
	background-color: #fff;	
	margin: 0 auto;
	max-width: 800px;
	padding: 30px;
}
.templatemo-payment-form .control-label { margin-bottom: 10px; }
.templatemo-payment-form .btn {	width: 150px; }
.templatemo-select-container { padding: 0; }
.cvv2 {
	width: 50px;
	padding: 5px;
	display: inline-block;
}
.cvv2-group > div {	display: inline-block; }
.cvv2-group > div > label {	display: block; }
.cvv2-group img { vertical-align: bottom; }
.templatemo-card-details label { font-weight: 400; }
.btn-round { border-radius: 30px; }
.templatemo-inline-group {
	display: inline-block;
	vertical-align: middle;
	height: 71px;
	padding-right: 20px;
}
.form-horizontal .control-label.text-left {	text-align: left; }
.radio-inline {	margin-right: 10px; }
.radio-inline+.radio-inline { margin-left: 0; }
.templatemo-radio-container { margin-left: 15px; }


/*-------------------------------------
10. Media Queries
---------------------------------------*/
@media screen and (max-width: 767px) {
	.templatemo-header .col-md-12 {	line-height: 20px; }
	.templatemo-header .btn { margin-bottom: 15px; }
	.templatemo-container {	padding: 15px; }
	.control-wrapper label.fa-label { top: 6px;	}
	.templatemo-header form { margin-top: 0; }
}
@media screen and (max-width: 991px) {
	.templatemo-login-form-2 .templatemo-one-signin {
		border-right: none;
		border-bottom: 1px solid rgba(200, 200, 200, 0.5);
	}
	.templatemo-login-form-2 .templatemo-other-signin {	padding-top: 30px; }
	.templatemo-login-form-2 .form-group:last-child { margin-bottom: 30px; }
	.templatemo-create-account .col-md-6, .templatemo-create-account .col-md-12 {
		padding-top: 5px;
		padding-bottom: 15px;
	}
	.templatemo-create-account .form-group { margin-bottom: 0; }
	.templatemo-radio-group { margin-top: 0; }
	.templatemo-contact-form-2 .checkbox {	margin-bottom: 15px; }
	.templatemo-contact-form-2 { padding: 30px 30px 0 30px;	}
	.templatemo-radio-container { margin-left: 30px; }
}
</style>