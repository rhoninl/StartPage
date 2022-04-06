let myFavourite
function showBox(url,width,height,parentWindow){
    layui.use('layer',function (){
        var layer = layui.layer;
        layer.open({
            title:'',
            btn: false,
            closeBtn:1,
            type: 2,
            area:[width,height],
            scrollbar:false,
            content: [url],
            end:function (){
                parentWindow.location.reload()
            }

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
            area: ['300px','500px'],
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
       showBox('Login','300px','200px',window);
       return
   }
    layui.use('layer',function () {
        var layer = layui.layer;
        layer.open({
            type: 2,
            title: false,
            area: ['300px','110px'],
            moveType: 1,
            content: ['/Setting'],
        });
    })
}

function judgeLogin() {
    let isLogin = document.getElementById("isLogin").innerText;
   if (isLogin[0] == "f") {
       document.getElementById("favourite-card").style.display= "none";
   }else{
       document.getElementById("login").onclick = LogOut;
       document.getElementById("login").innerText = "注销";
       document.getElementById('favourite-list').innerHTML = RenderFavourite(getFavourite().data);
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

function getFavourite(){
    let result
    $.ajax({
        type:"GET",
        url:"/Favourite",
        dataType:'json',
        async:false,
        success:function(data){
            result = data
            myFavourite = data.data
        },
        error:function(){
            alert("获取收藏夹失败");
        }
    });
    return result
}

function RenderFavourite(data){
    if (data == null) {
        document.getElementById("favourite-card").style.display= "none";
        return
    }
    let str = "<table class = 'favourite-table'>"
    for(let k in data){
        str += "<tr><td><img src='http://"+data[k]['Url']+"/favicon.ico'/><a href=http://"+data[k]['Url']+"><span>"+data[k]['Alias']+"</span></td></tr>"
    }
    str += "</table>"
    return str
}

function RenderFavouriteTable(data){
    let str = "<table class = 'favouriteTable'>"
    for(let k in data){
        str += "<tr><td><img src='http://"+data[k]['Url']+"/favicon.ico'/></td><td style='width: 150px'><span>"+data[k]['Alias']+"</span></td><td><button onclick='Alter(this)' value='"+k+"'>修改</button></td><td><button value='"+k+"' onclick='Delete(this)'>删除</button></td></tr>"
    }
    str += "<tr><td colspan='2'></td><td colspan='3'><button onclick='AddFavourite()'>添加</button></td></tr>"
    str += "</table>"
    return str
}

function AddFavourite(){
    parent.showBox("AddFavourite",'300px','100px',window)
}

function Delete(button){
    let data = myFavourite[button.value]
    console.log(data)
    $.ajax({
        type:"POST",
        url:"/DeleteFavourite",
        dataType:'json',
        contentType:"application/json",
        data:JSON.stringify({
            "FavouriteId":data['FavouriteId'],
        }),
        async:false,
        success:function(data){
            parent.layer.msg('删除成功');
            location.reload();
        },
        error:function(){
            parent.layer.msg('删除失败');
        }
    });
}

function Alter(button) {
    let layui = parent.layui
    layui.use('layer',function (){
        var layer = layui.layer;
        layer.open({
            title:'',
            btn: false,
            closeBtn:1,
            type: 2,
            area:['300px','120px'],
            scrollbar:false,
            content: ['/AlterFavourite/'+myFavourite[button.value]['FavouriteId']],
            end:function (){
                location.reload()
            },
        });
    })
}