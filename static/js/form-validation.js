let attention = Prompt();
        (() => {
            'use strict'

            // Fetch all the forms we want to apply custom Bootstrap validation styles to
            const forms = document.querySelectorAll('.needs-validation')

            // Loop over them and prevent submission
            Array.from(forms).forEach(form => {
                form.addEventListener('submit', event => {
                if (!form.checkValidity()) {
                    event.preventDefault()
                    event.stopPropagation()
                }

                form.classList.add('was-validated')
                }, false)
            })
        })()

        function Notify(msg, msgType) {
            notie.alert({
            type: msgType, // optional, default = 4, enum: [1, 2, 3, 4, 5, 'success', 'warning', 'error', 'info', 'neutral']
            text: msg,
            // stay: Boolean, // optional, default = false
            // time: Number, // optional, default = 3, minimum = 1,
            // position: String // optional, default = 'top', enum: ['top', 'bottom']
            })
        }
        function notifyModal(title, text, icon, confirmationButtonText) {
            Swal.fire({
            title: title,
            html: text,
            icon: icon,
            confirmButtonText: confirmationButtonText
            })
        }