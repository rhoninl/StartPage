
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
    layui.use('layer',function () {
        var layer = layui.layer;
        layer.open({
            type: 2,
            title: false,
            closeBtn: 2,
            area: ['300px','600px'],
            btn:['项目仓库'],
            btnAlign: 'c',
            moveType: 1,
            content: ['/Notice'],
            success: function(layero){
                var btn = layero.find('.layui-layer-btn');
                btn.find('.layui-layer-btn0').attr({
                    href:'https://github.com/MrLeea-13155bc/StartPage'
                })
            }
        });
    })
}

function Setting(){
   let isLogin = document.getElementById("isLogin").innerText;
   if(isLogin[0] == "f") {
       showBox('Login');
       return
   }
    layui.use('layer',function () {
        var layer = layui.layer;
        layer.open({
            type: 2,
            title: false,
            area: ['300px','100px'],
            moveType: 1,
            content: ['/Setting'],
        });
    })
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

function GetFavourite(){
    $.ajax({
        type:"POST",
        url:"/Favourite",
        success:function (result){
            return result.responseJSON[0]
        },
        error:function (data) {
            alert("获取收藏夹失败");
        }
    });
}

function ShowFavourite(){
    let favourites = GetFavourite();
    if (favourites.length == 0) {
        document.getElementById("favourite-card").style.display = 'none';
    }
}