// mii-temp-1.labs.novr.one

var xhr = new XMLHttpRequest();
var url = 'http://mii-temp-1.labs.novr.one/api/v1/check';
xhr.open('GET', url, true);
xhr.setRequestHeader('Authorization', 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJ1c2VybmFtZSI6ImFsbWlnaHR5dXNlciIsImlhdCI6NDA3ODMwMzA1Mn0.jj_bhYTrXkjSbpcxDNyY8Xq3Y_Oibx7xbY970zGE4RE');

xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
        // Request was successful, handle response
        var response = xhr.responseText;
        console.log(response); // Or do something else with the response
    } else {
        // Request failed or is still processing
        // You can handle other status codes or errors here
        console.log('Failed to get status');
    }
};

xhr.send();
