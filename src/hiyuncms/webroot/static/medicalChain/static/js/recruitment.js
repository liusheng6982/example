window.onload=function(){
    var one=$("#one");
    var two=$("#two");
    var three=$("#three");
    var oRaduis=$('div[name="radius"]');
    oRaduis[0].style.background = "#45b396";
    one.show();
    two.hide();
    three.hide();
    for(var i=0;i<oRaduis.length;i++){
        oRaduis[i].style.background="#e6e6e6";
    }
    var count=0;
    window.setInterval(function(){

           if (count % 3 == 0) {
               one.show();
               two.hide();
               three.hide();
               oRaduis[0].style.background = "#45b396";
               oRaduis[1].style.background = "#e6e6e6";
               oRaduis[2].style.background = "#e6e6e6";
           }
           if (count % 3 == 1) {
               one.hide();
               two.show();
               three.hide();
               oRaduis[0].style.background = "#e6e6e6";
               oRaduis[1].style.background = "#45b396";
               oRaduis[2].style.background = "#e6e6e6";
           }

           if (count % 3 ==2) {
               one.hide();
               two.hide();
               three.show();
               oRaduis[0].style.background = "#e6e6e6";
               oRaduis[1].style.background = "#e6e6e6";
               oRaduis[2].style.background = "#45b396";
           }
           if(count>=10){
               count=0;
           }
        count++;
    },5000)
    $("#radius1").click(function(){
        one.show();
        two.hide();
        three.hide();
        oRaduis[0].style.background = "#45b396";
        oRaduis[1].style.background = "#e6e6e6";
        oRaduis[2].style.background = "#e6e6e6";

    })
    $("#radius2").click(function(){
        one.hide();
        two.show();
        three.hide();
        oRaduis[0].style.background = "#e6e6e6";
        oRaduis[1].style.background = "#45b396";
        oRaduis[2].style.background = "#e6e6e6";

    })
    $("#radius3").click(function(){
        one.hide();
        two.hide();
        three.show();
        oRaduis[0].style.background = "#e6e6e6";
        oRaduis[1].style.background = "#e6e6e6";
        oRaduis[2].style.background = "#45b396";

    })

    var classitem = $('div[name="Item"]');
    for(var i = 0 ;i<classitem.length;i++){
        $(classitem[i]).click(function () {
          // alert("adf");
        })
    }
}