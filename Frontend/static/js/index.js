// Выпадающий список

jQuery(($) => {
    $('.select').on('click', '.select__head', function () {
        if ($(this).hasClass('open')) {
            $(this).removeClass('open');
            $(this).next().fadeOut();
        } else {
            $('.select__head').removeClass('open');
            $('.select__list').fadeOut();
            $(this).addClass('open');
            $(this).next().fadeIn();
        }
    });

    $('.select').on('click', '.select__item', function () {
        $('.select__head').removeClass('open');
        $(this).parent().fadeOut();
        $(this).parent().prev().text($(this).text());
        $(this).parent().prev().prev().val($(this).text());
    });

    $(document).click(function (e) {
        if (!$(e.target).closest('.select').length) {
            $('.select__head').removeClass('open');
            $('.select__list').fadeOut();
        }
    });
});

// /Выпадающий список

// Текст в инпуте с загрузкой файла

$("#file").change(function() {
  let i = $(this).prev('label').clone();
  let file = $('#file')[0].files[0].name;
  $(this).prev('label').text(file);
  $(this).prev('label').addClass("active");
});

// /Текст в инпуте с загрузкой файла

// Отправка формы

const myForm = document.querySelector("#form")
const myFormFile = document.querySelector("#file")
const myFormModel = document.querySelector(".select__input")
const message = document.querySelector(".message")

const apiBase = "http://192.168.0.111:8080/api/v1/new" // ссылка на файл
const timeOutMessage = 1000

myForm.onsubmit = function(e) {
	e.preventDefault()
	const formData = new FormData()
	if (!myFormModel.value) {
		message.innerHTML = "Выберите название машинки"
		message.classList.add("error", "active")
		setTimeout(() => {
			message.classList.remove("error", "active")
		}, timeOutMessage)
		return
	}
	if (!myFormFile.files[0]) {
		message.innerHTML = "Вы не выбрали файл"
		message.classList.add("error", "active")
		setTimeout(() => {
			message.classList.remove("error", "active")
		}, timeOutMessage)
		return
	}
	formData.append("name", myFormModel.value)
	formData.append("file", myFormFile.files[0])
	axios.post(apiBase, formData, {timeout: 2000})
	  .then(function (response) {
		  if (response.statusText == "OK") {
			  message.innerHTML = response.data.message
			  message.classList.add("active")
			  setTimeout(() => {
				  message.classList.remove("active")
			  }, timeOutMessage)
		  }
	  })
	  .catch(function (error) {
		  message.innerHTML = "Ошибка при отправке данных на сервер"
		  message.classList.add("error", "active")
		  setTimeout(() => {
			 message.classList.remove("error", "active")
		  }, timeOutMessage)
	  });
}

// /Отправка формы