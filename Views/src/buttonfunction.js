
function showBox(url){
    layui.use('layer',function (){
        var layer = layui.layer;
        layer.open({
            title:'',
            btn: false,
            closeBtn:1,
            type: 2,
            area:['300px','200px'],
            scrollbar:false,
            content: [url]
        });
    })
}

function About(){
    alert('尚在完善');
}

function Setting(){
    alert('尚在完善');
}

function judgeLogin() {
    var islogin = document.getElementById("isLogin").innerText;
    console.log(islogin)
   if (islogin[0] == "f") {
       document.getElementById("login").onclick = Login;
       document.getElementById("login").innerHTML = "登陆";
   }else{
       document.getElementById("login").onclick = LogOut;
       document.getElementById("login").innerText = "注销";
   }
}

function LogOut(){
    $.ajax({
        type:"POST",
        url:"/LogOut",
        success:function (result){
            parent.location.reload();
            parent.layer.msg('注销成功！');
        },
        error:function (data) {
            alert("LogOut Default!");
        }
    });
}