// Add Emp Model 
// document.addEventListener("DOMContentLoaded", () => {
//     const AddBtn = document.getElementById("addEmpModal")
//     if (AddBtn) {
//         AddBtn.addEventListener("click", () => {
//             const submitModal = document.getElementById("submit-add-form");
//             const form = document.getElementById("employee-form");

//             submitModal.addEventListener("click", function (event) {
//                 event.preventDefault();

//                 const formData = new FormData(form);
//                 const empName = formData.get('emp-name');
//                 const empDep = formData.get('emp-dep');
//                 const empDeg = formData.get('emp-deg');
//                 const empMob = formData.get('emp-mob');
//                 const empDOB = formData.get('emp-age');
//                 const empEmail = formData.get('emp-email');
//                 const empAddress = formData.get('emp-add');
//                 const empSalary = formData.get('emp-salary');
//                 const empDOJ = formData.get('emp-join');
//                 const empGendher = formData.get('emp-gender');
//                 const strempDOB = empDOB.toString();
//                 const strempDOJ = empDOJ.toString();
//                 const strempGendher = empGendher.toString();


//                 if (empMob.length !== 10) {
//                     let empMob = document.getElementById("emp-mob")
//                     alert('Please enter correct mobile number');
//                     empMob.focus();
//                     return;
//                 } else {
//                     // Send the data to the backend API using fetch
//                     fetch('/add-emp', {
//                         method: 'POST',
//                         body: JSON.stringify({
//                             name: empName,
//                             department: empDep,
//                             designation: empDeg,
//                             birthdate: strempDOB,
//                             joindate: strempDOJ,
//                             gender: strempGendher,
//                             email: empEmail,
//                             address: empAddress,
//                             mobile: empMob,
//                             salary: empSalary,

//                         }),
//                         headers: {
//                             'Content-Type': 'application/json'
//                         },
//                     }).then(response => {
//                         if (!response.ok) {
//                             alert("Fail to add employee details");
//                             return;
//                         }
//                         alert("Employee details added");
//                     })
//                 }
//                 window.location.href = "/dashboard";
//             });
//         })
//     }
// })



// delete function
document.addEventListener("DOMContentLoaded", () => {
    document.addEventListener("click", (event) => {
        const deleteBtn = event.target.closest(".delete-btn");
        if (deleteBtn) {
            event.preventDefault();

            const empId = deleteBtn.getAttribute("data-id");

            if (confirm("Are you sure you want to delete this employee?")) {
                fetch('/delete-emp', {
                    method: 'DELETE',
                    body: JSON.stringify({ id: empId }),
                    headers: {
                        'Content-Type': 'application/json'
                    }
                })
                    .then(response => {
                        if (response.ok) {
                            window.location.href = "/dashboard";
                        } else {
                            alert("Failed to delete employee. Please try again.");
                        }
                    })
                    .catch(error => {
                        console.error('Error deleting employee:', error);
                        alert("An error occurred while deleting the employee.");
                    });
            }
        }
    });
});


document.addEventListener("DOMContentLoaded", () => {
    // Handle Edit Button Click
    document.addEventListener("click", (event) => {
        const editBtn = event.target.closest(".edit-btn");
        if (editBtn) {
            const empId = editBtn.getAttribute("data-id");

            fetch('/EditModal', {
                method: 'POST',
                body: JSON.stringify({ id: empId }),
                headers: { 'Content-Type': 'application/json' }
            })
            .then(response => response.text()) // Parse response as text (HTML)
            .then(html => {
                document.getElementById('ModalHere').innerHTML = html;
                new bootstrap.Modal(document.getElementById('EditModal')).show();

                const closeButtons = document.getElementById("closeBtn")
                    closeButtons.addEventListener('click', () => {
                        new bootstrap.Modal(document.getElementById('EditModal')).hide();
                        document.getElementById('ModalHere').innerText= "";
                        window.location.href = "/dashboard";
                    });

                    const form = document.getElementById("EditEmpForm");
                    const SubmitBtn = document.getElementById("submit-form");
                    if (form) {
                        SubmitBtn.addEventListener("click", (event) => {
                            event.preventDefault();
                            
                            const empId = document.getElementById("id-for-emp").innerText;
                            const empName = document.getElementById('emp-name').value;
                            const empDep = document.getElementById('emp-dep').value;
                            const empDeg = document.getElementById('emp-deg').value;
                            const empMob = document.getElementById('emp-mob').value;
                            const empDOB = document.getElementById('emp-dob').value;
                            const empEmail = document.getElementById('emp-email').value;
                            const empAddress = document.getElementById('emp-add').value;
                            const empSalary = document.getElementById('emp-salary').value;
                            const empDOJ = document.getElementById('emp-join').value;
                            const empGender = document.getElementById('emp-gender').value;
                            
                            if (empMob.length !== 10) {
                                alert('Please enter a correct mobile number');
                                document.getElementById("emp-mob").focus();
                                return;
                            }
                
                            fetch('/update', {
                                method: 'PUT',
                                body: JSON.stringify({
                                    id: empId,
                                    name: empName,
                                    department: empDep,
                                    designation: empDeg,
                                    birthdate: empDOB,
                                    joindate: empDOJ,
                                    gender: empGender,
                                    email: empEmail,
                                    address: empAddress,
                                    mobile: empMob,
                                    salary: empSalary
                                }),
                                headers: { 'Content-Type': 'application/json' }
                            })
                            .then(response => {
                                if (!response.ok) {
                                    alert("Failed to update employee details");
                                    return;
                                }
                                alert("Employee details updated");
                                window.location.href = "/dashboard"; // Redirect to dashboard after success
                            })
                            .catch(error => {
                                console.error('Error updating employee:', error);
                                alert("An error occurred while updating the employee.");
                            });
                        });
                    }
                
            })
            .catch(error => {
                console.error("Error fetching modal:", error);
            });
        }
    });
});



// Add Emp modal
document.addEventListener("DOMContentLoaded", () => {
    // Handle Edit Button Click
    const AddBtn = document.getElementById("addEmpModal")
    if (AddBtn) {
        AddBtn.addEventListener("click", () => {
            fetch('/AddModal', { 
                method: 'GET',
                headers: { 'Content-Type': 'application/json' }
            })
            .then(response => response.text()) 
            .then(html => {
                document.getElementById('ModalHere').innerHTML = html;
                new bootstrap.Modal(document.getElementById('AddModal')).show();
                console.log(html); // Optional logging for debugging

                const closeButtons = document.getElementById("closeBtn")
                    closeButtons.addEventListener('click', () => {
                        new bootstrap.Modal(document.getElementById('AddModal')).hide();
                        document.getElementById('ModalHere').innerText= "";
                        window.location.href = "/dashboard";
                    });

                    const form = document.getElementById("AddEmpForm");
                    const SubmitBtn = document.getElementById("submit-form");
                    if (form) {
                        SubmitBtn.addEventListener("click", (event) => {
                            event.preventDefault();
                            
                            const empName = document.getElementById('emp-name').value;
                            const empDep = document.getElementById('emp-dep').value;
                            const empDeg = document.getElementById('emp-deg').value;
                            const empMob = document.getElementById('emp-mob').value;
                            const empDOB = document.getElementById('emp-dob').value;
                            const empEmail = document.getElementById('emp-email').value;
                            const empAddress = document.getElementById('emp-add').value;
                            const empSalary = document.getElementById('emp-salary').value;
                            const empDOJ = document.getElementById('emp-join').value;
                            const empGender = document.getElementById('emp-gender').value;
                            
                            if (empMob.length !== 10) {
                                alert('Please enter a correct mobile number');
                                document.getElementById("emp-mob").focus();
                                return;
                            }
                
                            fetch('/add-emp', {
                                method: 'POST',
                                body: JSON.stringify({
                                    name: empName,
                                    department: empDep,
                                    designation: empDeg,
                                    birthdate: empDOB,
                                    joindate: empDOJ,
                                    gender: empGender,
                                    email: empEmail,
                                    address: empAddress,
                                    mobile: empMob,
                                    salary: empSalary
                                }),
                                headers: { 'Content-Type': 'application/json' }
                            })
                            .then(response => {
                                if (!response.ok) {
                                    alert("Failed to add employee details");
                                    return;
                                }
                                alert("Employee Details Added");
                                window.location.href = "/dashboard"; // Redirect to dashboard after success
                            })
                            .catch(error => {
                                console.error('Error adding employee:', error);
                                alert("An error occurred while updating the employee.");
                            });
                        });
                    }
                
            })
            .catch(error => {
                console.error("Error fetching modal:", error);
            });
        });
    }
});

// All Delete Button
document.addEventListener("DOMContentLoaded", () => {
    const checkBox = document.getElementById("allCheckBox");
    const deleteButton = document.getElementById("allDeleteButton");
    const rowCheckboxes = document.querySelectorAll('tbody .form-check-input');

    // Function to update the delete button's text and log emails of selected rows
    const updateDeleteButton = () => {
        // Check if all row checkboxes are checked
        const allChecked = [...rowCheckboxes].every(checkbox => checkbox.checked);

        // Check if either allCheckBox or any row checkbox is checked
        const anyChecked = checkBox.checked || [...rowCheckboxes].some(checkbox => checkbox.checked);

        if (anyChecked) {
            if (deleteButton) {
                deleteButton.innerHTML = "Delete Karde bhai sub";

                // Extract emails of selected rows
                const selectedEmails = [...rowCheckboxes]
                    .filter(checkbox => checkbox.checked)  // Get only checked checkboxes
                    .map(checkbox => checkbox.closest('tr').querySelector('.Email').innerText);  // Extract the email from the same row

                // Log the list of selected emails
                console.log("Selected emails:", selectedEmails);
                deleteButton.addEventListener("click", function () {
                    deleteButton.innerHTML = "Ho Gaya sub delete bhai";
                });
            }
        } else {
            if (deleteButton) {
                deleteButton.innerHTML = "";
            }
        }

        // Sync the state of #allCheckBox with the state of the row checkboxes
        checkBox.checked = allChecked;
    };

    // Make sure the checkBox and deleteButton exist
    if (checkBox && deleteButton) {
        // Add an 'onchange' event listener to the #allCheckBox
        checkBox.addEventListener("change", function () {
            const isChecked = checkBox.checked;
            // Set all row checkboxes to the same state as #allCheckBox
            rowCheckboxes.forEach(checkbox => {
                checkbox.checked = isChecked;
            });
            updateDeleteButton();
        });

        // Add 'change' event listeners to each row checkbox
        rowCheckboxes.forEach(checkbox => {
            checkbox.addEventListener("change", function () {
                updateDeleteButton();
            });
        });
    }
});