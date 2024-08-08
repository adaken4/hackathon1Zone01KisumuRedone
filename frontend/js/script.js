document.getElementById('signupForm').addEventListener('submit', function(event) {
  event.preventDefault(); // Prevent the form from submitting

  var password = document.getElementById('password').value;
  var confirmPassword = document.getElementById('confirmPassword').value;

  if (password !== confirmPassword) {
    document.getElementById('errorMessage').textContent = 'Passwords do not match.';
  } else {
    // Passwords match, submit the form
    this.submit();
  }
});