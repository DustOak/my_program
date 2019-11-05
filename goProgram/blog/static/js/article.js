$(document).ready( function () {
    $('#myTable').DataTable({
        "autoWidth": true,
        "info":true,
        "lengthChange": true,
        "searching": false,
        "stateSave": true,
        "ordering":false,
        "orderClasses": false,
        "pagingType": "full_numbers",
        "language": {
            "lengthMenu": "每页 _MENU_ 条记录",
            "zeroRecords": "没有找到记录",
            "info": "第 _PAGE_ 页 ( 总共 _PAGES_ 页 )",
            "infoEmpty": "无记录",
            "infoFiltered": "(从 _MAX_ 条记录过滤)",
            "oPaginate": {
                "sFirst": "首页",
                "sPrevious": "上页",
                "sNext": "下页",
                "sLast": "末页"
            },
        },
    });
} );

function delArticle(id) {
   if (confirm("确认删除?")){
       $.ajax({
           url:"/admin/article/delete",
           method:"POST",
           data:{
               id:id
           },
           success:function (result) {
                if (result=="-1"){
                    alert("删除失败");
                }else {
                    alert("删除成功");
                    location.reload();
                }
           }
       })
   }
};

function delComment(id) {
    alert(id)
};

function disp() {
    document.getElementById("addArticle").style.display="block";
}
function cls() {
    document.getElementById("addArticle").style.display="none";
}
