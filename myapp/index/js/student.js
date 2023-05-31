function addStudent() {
    var data = getFormData()

    fetch('/student', {
        method: "POST",
        body: JSON.stringify(data),
        headers: {"Content-type": "application/json: charset=UTF-8"}
    }).then (response1 => {
        if (response1.ok) {
            fetch('/student/'+sid)
            .then(response2 => response2.text())
            .then(data => showStudent(data))
        } else {
            throw new Error(response1.statusText)
        }
    }).catch(e => {
        alert(e)
    })
    resetform();
}


window.onload = function () {
    fetch("/students")
    .then(response => response.text())
    .then(data => showStudents(data))
}

function showStudent(data) {
    const student = JSON.parse(data)
    newRow(student)
}

function showStudents(data) {
    const students = JSON.parse(data)
    students.forEach(stud => {
        newRow(stud)
    });
}


//set form fields to empty
function resetform() {
    document.getElementById("sid").value = "";
    document.getElementById("fname").value = "";
    document.getElementById("lname").value = "";
    document.getElementById("email").value = "";
}


function newRow(student) {
    //Find a <table> element with id="myTable":
    var table = document.getElementById("myTable");

    //create an empty <tr> element and add to the last position of the table:
    var row = table.insertRow(table.length)

    //insert new cells (<td>elements) at the 1st and 2nd position of the "new" <tr> element:
    var td = []
    for (i=0; i<table.rows[0].cells.length; i++) {
        td[i] = row.insertCell(i);
    }

    //Add student detail to the new cells:
    td[0].innerHTML = student.stdid;
    td[1].innerHTML = student.fname;
    td[2].innerHTML = student.lname;
    td[3].innerHTML = student.email;
    td[4].innerHTML = '<input type="button" onclick="deleteStudent(this)" value="delete" id="button-1">';
    td[5].innerHTML = '<input type="button" onclick="updateStudent(this)" value="edit" id="button-2">';
}

var selectedRow = null;

function updateStudent(r) {
    selectedRow = r.parentElement.parentElement;

    //fill in the form fields with the selected row data
    document.getElementById("sid").value = selectedRow.cells[0].innerHTML;
    document.getElementById("fname").value = selectedRow.cells[1].innerHTML;
    document.getElementById("lname").value = selectedRow.cells[2].innerHTML;
    document.getElementById("email").value = selectedRow.cells[3].innerHTML;

    var btn = document.getElementById("button-add")
    sid = selectedRow.cells[0].innerHTML;
    if (btn) {
        btn.innerHTML = "Update";
        btn.setAttribute("onclick", "update(sid)");
    }
}

function getFormData() {
    var formData = {
        stdid : parseInt(document.getElementById("sid").value),
        fname : document.getElementById("fname").value,
        lname : document.getElementById("lname").value,
        email : document.getElementById("email").value
    }
    return formData
}

function update(sid) {
    //data to be sent to the UPDATE request
    var newData = getFormData()

    fetch('/student/'+sid, {
        method: "PUT",
        body: JSON.stringify(newData),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then (res => {
        if (res.ok) {
            //fill in selected row with updated value
            selectedRow.cells[0].innerHTML = newData.stdid;
            selectedRow.cells[1].innerHTML = newData.fname;
            selectedRow.cells[2].innerHTML = newData.lname;
            selectedRow.cells[3].innerHTML = newData.email;
            //set to previous value
            var button = document.getElementById("button-add");
            button.innerHTML = "Add";
            button.setAttribute("onclick", "addStudent()");
            selectedRow = null;

            resetform();
        } else {
            alert("Server: Update request error.")
        }
    })
}

function deleteStudent(r){
    if (confirm('Are you sure you want to DELETE this?')){
        selectedRow = r.parentElement.parentElement;
        sid = selectedRow.cells[0].innerHTML;

        fetch('/student/'+sid, {
            method: "DELETE",
            headers: {"Content-type": "application/json; charset=UTF-8"}
        });
        var rowIndex = selectedRow.rowIndex;
        if (rowIndex>0){
            document.getElementById("myTable").deleteRow(rowIndex);
        }
        selectedRow = null;
    }
}
