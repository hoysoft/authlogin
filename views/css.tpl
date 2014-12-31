<style>
.login-input {
	font-size: 25px;
	margin-bottom: 15px;
	border-color: transparent;
	font-family: "Hevetical Neue", Helvetica, Arial, sans;
	font-weight: 200
}

.login-form {
	-webkit-animation: login-form-fade-in 0.6s cubic-bezier(0.19, 1, 0.98, 1);
	animation: login-form-fade-in 0.6s cubic-bezier(0.19, 1, 0.98, 1)
}

.login-footer {
	margin: 30px 0;
	text-align: center;
	color: #5c626b
}

.login-footer a {
	color: #8f959e;
	margin-bottom: 1em;
	text-decoration: none;
	font-family: "Hevetical Neue", Helvetica, Arial, sans;
	font-weight: 200
}

.login-footer a:hover {
	color: #aaafb6;
	text-decoration: underline
}


.login-page {
	max-width: 400px;
	margin: 0 auto;
	padding: 0 1em;
	position: relative
}

 .triangle-obtuse {
  position:relative;
  padding:15px;
 /* margin:1em 0 3em; */
  color:#fff;
  background:#c81e2b;
  /* css3 */
  background:-webkit-gradient(linear, 0 0, 0 100%, from(#f04349), to(#c81e2b));
  background:-moz-linear-gradient(#f04349, #c81e2b);
  background:-o-linear-gradient(#f04349, #c81e2b);
  background:linear-gradient(#f04349, #c81e2b);
  -webkit-border-radius:10px;
  -moz-border-radius:10px;
  border-radius:10px;
}

	.triangle-border {
  position:relative;
  padding:2px;
/*  margin:1em 0 3em; */
  border:3px solid  #c81e2b;
 
  /* #5a8f00; */
  color:#333;
  background:#fff;
  /* css3 */
  -webkit-border-radius:5px;
  -moz-border-radius:5px;
  border-radius:5px;
}


/* THE TRIANGLE
------------------------------------------------------------------------------------------------------------------------------- */

.triangle-border:before {
  content:"";
  position:absolute;
  bottom:-10px; /* value = - border-top-width - border-bottom-width */
  left:40px; /* controls horizontal position */
  border-width:10px 10px 0;
  border-style:solid;
  border-color:#c81e2b transparent;
 /* border-color:#5a8f00 transparent;*/
  /* reduce the damage in FF3.0 */
  display:block;
  width:0;
}

/* creates the smaller  triangle */
.triangle-border:after {
  content:"";
  position:absolute;
  bottom:-3px; /* value = - border-top-width - border-bottom-width */
  left:47px; /* value = (:before left) + (:before border-left) - (:after border-left) */
  border-width:3px 13px 0;
  border-style:solid;
  border-color:#fff transparent;
  /* reduce the damage in FF3.0 */
  display:block;
  width:0;
}


	/* creates the larger triangle */
.triangle-border.top:before {
  top:-10px; /* value = - border-top-width - border-bottom-width */
  bottom:auto;
  left:auto;
  right:40px; /* controls horizontal position */
  border-width:0 10px 10px;
}

/* creates the smaller  triangle */
.triangle-border.top:after {
  top:-3px; /* value = - border-top-width - border-bottom-width */
  bottom:auto;
  left:auto;
  right:47px; /* value = (:before right) + (:before border-right) - (:after border-right) */
 border-width:0 3px 3px;
}
</style>