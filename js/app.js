// Define constants
const cameraView = document.getElementById('camera--view'),
      cameraOutput = document.getElementById('camera--output'),
      cameraSensor = document.getElementById('camera--sensor'),
      cameraTrigger = document.getElementById('camera--trigger'),
      cameraToggle = document.getElementById('camera--toggle');

var track = null;
var cameraNum = 0;
var cameraArray = ["user","environment"];

// Set constraints for the video stream
//2560*1440为1080P的分辨率
var constraints = { audio: false,
    video: {
        width: { min: 320, ideal: 1280, max: 2560 },
        height: { min: 240, ideal: 720, max: 1440 },
        frameRate: {
            ideal: 120,
            min: 10
        },
        facingMode: "user"       
    }
};

// Access the device camera and stream to cameraView
function cameraStart() {
    navigator.mediaDevices
        .getUserMedia(constraints)
        .then(function(stream) {
            track = stream.getTracks()[0];
            cameraView.srcObject = stream;
        })
        .catch(function(error) {
            console.error("Something is wrong.", error);
        });
};

//Toggle the camera when cameraToggle is tapped
cameraToggle.onclick = function() {
    cameraNum += 1;
    var camerafacingMode = cameraArray[cameraNum % cameraArray.length];
    constraints.video.facingMode = camerafacingMode;
    track.stop();
    cameraStart(constraints);
    console.log("toggle to ",camerafacingMode);    
};



// Take a picture when cameraTrigger is tapped
cameraTrigger.onclick = function() {
    cameraSensor.width = cameraView.videoWidth;
    cameraSensor.height = cameraView.videoHeight;
    cameraSensor.getContext("2d").drawImage(cameraView, 0, 0);
    cameraOutput.src = cameraSensor.toDataURL("image/png");
    cameraOutput.classList.add("taken");
    console.log("current is ",constraints.video.facingMode);
};

// Start the video stream when the window loads
window.addEventListener("load", cameraStart, false);
