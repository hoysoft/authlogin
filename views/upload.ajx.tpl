 <form method="post" id="uploadform"  action="{{.action}}" enctype="multipart/form-data">
 <center>
<div id="_filelists" style=" display:block;">
   <div class="input-group" id="filegroup_0"  fid="0" style="width:80%;padding:0.5em 0;">
     <input type="text" class="form-control" onclick="browse(this)" id="f_file_0" readonly="readonly" style="cursor:pointer;"/>
     <input type="file" name="file[]" class="hidefile" id="hidefile_0" onchange="onChange(this)"/> 
      <span class="input-group-btn">
         <button class="button btn btn-success" onclick="browse(this)"   type="button"><li class="glyphicon glyphicon-folder-open"/></button>
		 <button class="btn btn-danger disabled" type="button" onclick="delitem(this)"><li class="glyphicon glyphicon-trash"/></button>
      </span>
    </div> 
	 
 </div>	 

<div id="progress"  style="padding:0; display:none;" >
  <hr id="_hr" style="padding:0;display:none;">
  <p id="statu" style="font-size:12px;">xxx</p>

  <div class="progress"  id="progressv" style="width:80%;height:18px;padding:0; "  >
     <div class="progress-bar progress-bar-striped active" role="progressbar" id="progressbar" aria-valuenow="60" aria-valuemin="0" aria-valuemax="100" style="width: 60%;">
       60%
     </div>
  </div>
</div>

 </center>

 <div class="modal-footer" id="_foot">
     <button type="button" id="_add" class="btn btn-warning" onclick='additem();' disabled="true" ><li class="glyphicon glyphicon-plus"/>&nbsp;增加文件</button>
	 &nbsp;&nbsp;&nbsp;&nbsp;
 
	 <button type="submit" class="btn btn-primary"   id="_upload" disabled="true">提交</button>
     <button type="button" class="btn btn-default"   id="_close" data-dismiss="modal" >关闭</button>

 </div>  
<form>


<style>
.hidefile
{ 
width: 0px; 
height: 0px; 
display: none; 
} 
 
</style>

<script type="text/javascript"> 

 
//增加一个文件项目
  function additem() {		
	//计算id
 	var maxId=0;

	 $('.input-group').each(function(){ 
	   DragId =parseInt(this.getAttribute('fid',2));
	 //  alert(DragId);
	   if (DragId>=maxId) {
		 maxId=DragId+1;
	   };
	 });
   // alert(maxId);
	
	//////
	data = [];
	data.push('<div class="input-group" id="filegroup_'+maxId+'"  fid="'+maxId+'" style="width:80%;padding:0.5em 0;">');
    data.push(' <input type="text" class="form-control" onclick="browse(this)"  id="f_file_'+maxId+'" readonly="readonly"  style="cursor:pointer;"/>');
    data.push(' <input type="file" name="file[]" class="hidefile" id="hidefile_'+maxId+'" onchange="onChange(this)"/> ');
    data.push('  <span class="input-group-btn">');
    data.push('     <button class="button btn btn-success" onclick="browse(this)"  type="button"><li class="glyphicon glyphicon-folder-open"/></button>');
	data.push('	 <button class="btn btn-danger" type="button"  onclick="delitem(this)"><li class="glyphicon glyphicon-trash"/></button>');
    data.push('  </span>');
    data.push(' </div>');
	document.getElementById("_add").disabled="true";
     $("#_filelists").append(data.join(""));
	};
 
//浏览选择文件
 function browse(v){
 	var obj =getRootObj(v);
 	var hidefile="#hidefile_"+obj.getAttribute('fid',2);
    $(hidefile).click(); 
 };

//删除一个文件列表项目
function delitem(v){
	var obj =getRootObj(v);
	obj.parentNode.removeChild(obj);
	UpdateBtnSuatus();
};
 
//获取文件根节点
function getRootObj(v){
	var obj =v;
    while (obj.parentNode.className !="input-group") {
		obj=obj.parentNode;
	};
    obj=obj.parentNode;
	//alert(obj.className);
	//alert(obj.id);
    return obj;
};

//修改按钮状态
function UpdateBtnSuatus(){
    var total_size=0;
	document.getElementById("_add").removeAttribute("disabled");
	$('.hidefile').each(function(){ 
	   total_size=total_size+(this.value.length);
	   //如果存在空的文件选择控件，不允许增加
	   if (this.value.length==0) {
		 document.getElementById("_add").disabled="true";
	   };
	});
 
    if 	(total_size>0)	{
		  document.getElementById("_upload").removeAttribute("disabled");
	}else{
		  document.getElementById("_upload").disabled="true";
	};
}

//显示文本框文件名
function onChange(v){
	$('.hidefile').each(function(){ 
	   if (v!=this && this.value==v.value) { 
	     alert("不能选择重复的文件！");
		 this.value="";
		 this.reset();//清除
	     exit;
		}
	});
	
	 document.getElementById("progress").style.display="none"; 
	var obj =getRootObj(v);
	document.getElementById("f_file_"+obj.getAttribute('fid',2)).value=v.value
    UpdateBtnSuatus();
}

function setPercenta(progressbar,v){
	progressbar.style.width=v;
    $("#progressbar").html(v);
}

$(function() {  
        $("#uploadform").submit(function(){  
		   //如果存在空的文件选择控件，删除
		    $('.hidefile').each(function(){
				if (this.value.length==0) {
		             delitem(this);
	           };
			 });
		    progressbar=document.getElementById("progressbar");
            $(this).ajaxSubmit({  
              beforeSubmit: function(arr, $form, options) {
					   setPercenta(progressbar,"0%");
					   document.getElementById("_hr").style.display="none";
					   document.getElementById("_filelists").style.display="none";
					   document.getElementById("progress").style.display="block"; 
					   document.getElementById("progressv").style.display="block";
					   document.getElementById("_foot").style.display="none";
					   $("#statu").html("初始化...");
		            },
	          uploadProgress:function(data) {
				    if (data!=null) {
					  $("#statu").html("正在上传("+my_formatSize(data.position)+"/"+my_formatSize(data.total)+")...");
				 	  setPercenta(progressbar,Math.round(data.position / data.total  * 10000) / 100.00 + "%");
					}
			        },
	          error: function(request) {
					document.getElementById("_hr").style.display="block";
					document.getElementById("_filelists").style.display="block";
					document.getElementById("progressv").style.display="none";
				    document.getElementById("_foot").style.display="block";
					$("#statu").html("上传失败."+request);
                    },
              success: function(data) {
                   // $("#commonLayout_appcreshi").parent().html(data);
				   // alert("success");
				     $("#statu").html("上传成功！！！");
					// document.getElementById("_close").click();
                    }
            });  
            return false; //不刷新页面  
        });  
});  

	
 //字节大小单位转换
function my_formatSize($size){
    var size  = parseFloat($size);
    var rank =0;
    var rankchar ='Bytes';
    while(size>1024){
size = size/1024;
rank++;
    }
    if(rank==1){
rankchar="KB";
    }
    else if(rank==2){
rankchar="MB";
    }
    else if(rank==3){
rankchar="GB";
    }    
    return size.toFixed(2)+ " "+ rankchar;
}


</script> 