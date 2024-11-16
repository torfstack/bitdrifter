const canvasWidth = 1200;
const canvasHeight = 900;

const gameCanvas = {
    canvas: document.createElement("canvas"),
    start: function() {
        this.canvas.width = canvasWidth;
        this.canvas.height = canvasHeight;
        this.context = this.canvas.getContext("2d");
        this.context.strokeStyle = "black";
        document.getElementById("canvas").appendChild(this.canvas);
    },
};

function startGame() {
    console.log("Starting game...");
    gameCanvas.start();
    gameCanvas.context.save();
    drawSomething();
    rotate();
}

const rectX = 200;
const rectY = 200;
const rectWidth = 150;
const rectHeight = 100;

function drawSomething() {
    gameCanvas.context.strokeRect(rectX, rectY, rectWidth, rectHeight);
}

const fps = 144;

function rotate() {
    console.log("Rotating...");
    let rotation = 0;
    setInterval(function() {
        gameCanvas.context.reset();
        gameCanvas.context.translate(rectX+rectWidth/2,rectY+rectHeight/2);
        gameCanvas.context.rotate(rotation * Math.PI / 180);
        gameCanvas.context.translate(-rectX-rectWidth/2, -rectY-rectHeight/2);
        gameCanvas.context.rect(rectX, rectY, rectWidth, rectHeight);
        gameCanvas.context.stroke();
        rotation += 30/fps;
    }, 1000/fps)
}
