const canvas = document.querySelector('canvas');
const c = canvas.getContext('2d');

//  SETS THE HEIGHT AND WIDTH OF THE CANVAS 
canvas.width = window.innerWidth;
canvas.height= window.innerWidth;

//resize canvas with the window
window.addEventListener('resize', function(){
    canvas.width = window.innerWidth;
    canvas.height= window.innerWidth;
})


class Circle{
    constructor(x, y, xv, yv, radius,alpha) {
        // OBJECTS ATRIBUTTES
        this.x = x;
        this.y = y;
        this.xv = xv;
        this.yv = yv;
        this.alpha = alpha;
        this.radius = radius;
        this.color = colorArray[Math.floor(Math.random() * colorArray.length)] //Gets a random color on the COLOR ARRAY
        // OBJECT METHODS
        this.draw = function () {
            
            c.globalAlpha = this.alpha;
            c.beginPath();
            c.arc(this.x, this.y, this.radius, 0, Math.PI * 2, false);
            c.strokeStyle = this.color;
            c.stroke();
            
        };
        
        // UPDATES SOMETHING ON THE CIRCLE( VELOCITY, OPACITY, ETC..)
        this.update = function () {
            
            //  CHANGES THE "X" DIRECTION OF THE CIRCLE WHEN HITTING A WALL
            if (this.x + this.radius > innerWidth || this.x - this.radius < 0) {
                this.xv = -this.xv;
            }

            if( this.y > innerHeight || this.y == innerHeight){
                this.y = innerHeight - radius
            }

            
            if (this.y + this.radius > innerHeight) {
                this.yv = -this.yv;
            }
            //MAKES THE CIRCLES SLOWLY DISAPPEAR AFTER REACHING CERTAIN HEIGHT
            
            if (this.y +this.radius < innerHeight/1.68){
                this.alpha -= 0.01;
            }
            if (this.alpha < 0 ) {
                this.x = Math.random() * (innerWidth - this.radius * 2) + this.radius;
                this.y = innerHeight - this.radius;
                this.alpha = 1
            } 
            
            this.x += this.xv;
            this.y += this.yv;

            // LOWS THE OPACITY AFTER CROSSING HALF THE CANVAS
            
            this.draw();
            
        };
    }
}



// VARIABLES

// CIRCLE COLORS 
var circleArray = [];
var colorArray = [
    '#0F181C',
    '#010508',
    //'#0B447D',
];


circleArray = [];

for (var i = 0; i < 1500; i++) { 
var radius = Math.floor(Math.random() * 25);
var x = Math.random() * (innerWidth - radius * 2) + radius;
var y = innerHeight/1.68 + Math.random() * innerHeight/1.68; 
var xv = Math.random() -0.5;
var yv = Math.random() -1.5;
var alpha = 1 ;
circleArray.push( new Circle(x,y,xv,yv,radius,alpha));
}


//  ANIMATION LOOP  
function animate () {
    requestAnimationFrame(animate); // loops to create frames
    c.clearRect(0, 0, innerWidth, innerHeight); // clears the page(frame)
    //BACKGROUND
    c.globalAlpha = 1
    c.fillStyle = "#07090A"; // BG COLOR
    c.fillRect(0,0, canvas.width,canvas.height); //  FILLS THE CANVAS WITH A COLOR
    
    //updates every circle in the array
    for(var i = 0; i < circleArray.length;i++){
        circleArray[i].update();
    }
    console.log('frame');
}

animate();