var currentIndex = 0; //当前索引
var repeater = null;//定义一个重复器，用来滚动
var timer = null;//定义定时器用来实现自动滚动
window.onload=function(){
    move(1);
}
function move(index){
    var container = document.getElementById("container");
    var left = currentIndex * 255; // 当前的scrollLeft的值
    var end = index * 255; //想要移动到的scrollLeft的值
    var maxStep = 255;
    var step = (end - left) / maxStep;
    var currentStep = 0;//进行到多少步了
    repeater = setInterval(
        function(){
            //滚动操作
            if(currentStep == maxStep){
                //停止滚动
                clearInterval(repeater);
                //当前索引发生变化,变为已经滚动到的位置
                currentIndex = index;
                //为了实现自动滚动，可以定义一个定时器
                timer = setTimeout(
                    function(){
                        move(currentIndex + 1 == 3 ? 0 : currentIndex + 1);
                    },
                    1000
                );
            }
            currentStep++;
            container.scrollLeft += step;
        },1
    );
}

