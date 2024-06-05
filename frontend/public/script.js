document.addEventListener("DOMContentLoaded", function() {
    fetchEmployees();
});

function fetchEmployees() {
    fetch('/employees')
        .then(response => response.json())
        .then(data => {
            const employeeList = document.getElementById('employee-list');
            employeeList.innerHTML = '';
            data.forEach(employee => {
                const li = document.createElement('li');
                li.textContent = `${employee.name} - ${new Date(employee.birthDate).toLocaleDateString()}`;
                employeeList.appendChild(li);
            });
        });
}

function subscribe() {
    const userId = 1; // Для простоты используем статический userId
    const employeeId = document.getElementById('employee-id').value;
    fetch('/subscribe', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ userId, employeeId: parseInt(employeeId) }),
    });
}

function unsubscribe() {
    const userId = 1; // Для простоты используем статический userId
    const employeeId = document.getElementById('employee-id').value;
    fetch('/unsubscribe', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ userId, employeeId: parseInt(employeeId) }),
    });
}

function setNotificationTime() {
    const userId = 1; // Для простоты используем статический userId
    const time = document.getElementById('notification-time').value;
    fetch(`/users/${userId}/notification-time`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ time }),
    });
}
