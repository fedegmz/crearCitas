document.addEventListener('DOMContentLoaded', function () {

  var calendarEl = document.getElementById('calendar');
  let fecha = new Date();

  let year = fecha.getFullYear();
  let month = (fecha.getMonth() + 1).toString().padStart(2, '0'); // Sumar 1 al mes ya que los meses comienzan en 0
  let day = fecha.getDate().toString().padStart(2, '0');

  let fechaActual = `${year}-${month}-${day}`;

  var calendar = new FullCalendar.Calendar(calendarEl, {
    headerToolbar: {
      left: 'prev,next today myCustomButton',
      center: 'title',
      right: 'dayGridMonth,timeGridWeek,timeGridDay,listMonth'
    },
    initialDate: fechaActual,
    initialView: 'dayGridMonth',
    // locale: 'es',
    navLinks: true, // can click day/week names to navigate views
    businessHours: true, // display business hours
    editable: true,
    selectable: true,
    selectMirror: true,
    timeZone: 'America/Mexico_City',
    eventDrop: function (info) {

      fetch('http://localhost:3000/api/citas/' + info.event.id, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          usuario: "luca",
          fecha_inicio: formatDate(info.event.start.toISOString()),
          fecha_fin: formatDate(info.event.end.toISOString()),
        })
      })
        .then(response => response.json())
        .then(data => {
          console.log(data)
          if (data.success) {
            Swal.fire({
              icon: 'success',
              title: 'Cita actualizada correctamente',
              showConfirmButton: false,
              timer: 1500
            })

          } else {
            Swal.fire({
              icon: 'error',
              title: 'Oops...',
              text: 'No se pudo actualizar!',
            })
            info.revert();
          }
        })
        .catch((error) => {
          console.error('Error:', error);
          //delete event
          info.revert();
          //
        })
    }

    ,
    select: function (arg) {
      // var title = prompt('BIENVENIDO Title:');
      // if (title) {
      //   // Formatea las fechas en formato ISO 8601 con zona horaria 'America/Mexico_City'
      //   var startFormatted = arg.start.toISOString();
      //   var endFormatted = arg.end.toISOString();

      //   fetch('http://localhost:3000/api/citas', {
      //     method: 'POST',
      //     headers: {
      //       'Content-Type': 'application/json'
      //     },
      //     body: JSON.stringify({
      //       titulo: title,
      //       usuario: "luca",
      //       fecha_inicio: formatDate(startFormatted),
      //       fecha_fin: formatDate(endFormatted),
      //     })
      //   })
      //     .then(response => response.json())
      //     .then(data => {
      //      if(data.success){
      //       calendar.addEvent({
      //         title: title,
      //         start: startFormatted,
      //         end: endFormatted,
      //         allDay: arg.allDay
      //       })
      //       }else{
      //         console.log(data)
      //         alert('No se pudo crear la cita')

      //       }
      //     }
      //     )
      //     .catch((error) => {
      //       console.error('Error:', error);

      //     })
      // }
      let startFormatted = arg.start.toISOString();
      let endFormatted = arg.end.toISOString();
      console.log(startFormatted)
      console.log(endFormatted)
      Swal.fire({
        title: 'Ingresa el nombre de la cita',
        input: 'text',
        inputAttributes: {
          autocapitalize: 'off'
        },
        showCancelButton: true,
        cancelButtonText: 'Cancelar',
        confirmButtonText: 'Aceptar',
        showLoaderOnConfirm: true,
        preConfirm: (title) => {
          return fetch(`http://localhost:3000/api/citas`, {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json'
            },
            body: JSON.stringify({
              titulo: title,
              usuario: "luca",
              fecha_inicio: formatDate(startFormatted),
              fecha_fin: formatDate(endFormatted),
            })
          })
            .then(response => {
              console.log(title)
              if (!response.ok) {
                throw new Error(response.statusText)
              }
              return response.json().then(data => {
                data['title'] = title;
                return data;
              })
            })
            .catch(error => {
              Swal.showValidationMessage(
                `Request failed: ${error}`
              )
            })
        },
        allowOutsideClick: () => !Swal.isLoading()
      }).then((data) => {
        console.log(data)
        if (data.value.success) {
          Swal.fire({
            icon: 'success',
            title: 'Cita creada correctamente',
            showConfirmButton: false,
            timer: 1500
          })
          calendar.addEvent({
            title: data.value.title,
            start: startFormatted,
            end: endFormatted,
            allDay: arg.allDay
          })
        } else {
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: data.value.message,
          })
          throw new Error(data.statusText)
        }
      })

      calendar.unselect()

    },
    customButtons: {
      myCustomButton: {
        text: 'custom!',
        click: function () {
          alert('clicked the custom button!');
        }
      }
    },
    eventClick: function (arg) {
      // console.log(arg.event.id)
      // if (confirm('Are you sure you want to delete this event?')) {
      //   arg.event.remove()
      // }
      Swal.fire({
        title: `¿Estás seguro de eliminar la cita ${arg.event.title}?`,
        text: "No podrás revertir esto!",
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        confirmButtonText: '<i class="fas fa-trash-alt"></i> Eliminar',
        cancelButtonText: '<i class="fas fa-times"></i> Cancelar',
      }).then((result) => {
        if (result.isConfirmed) {

          try {
            fetch('http://localhost:3000/api/citas/' + arg.event.id, {
              method: 'DELETE',
              headers: {
                'Content-Type': 'application/json'
              }
            })
              .then(response => {
                if (!response.ok) {
                  throw new Error(response.statusText)
                }
                return response.json();
              })
              .then(data => {
                console.log(data)
                if (data.success) {
                  Swal.fire({
                    icon: 'success',
                    title: 'Cita eliminada correctamente',
                    showConfirmButton: false,
                    timer: 1500
                  })
                  arg.event.remove()
                } else {
                  Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: 'No se pudo eliminar!',
                  })
                }
              })
          } catch (error) {
            console.error('Error:', error);
            Swal.fire({
              icon: 'error',
              title: 'Oops..',
              text: 'Ocurrió un error al conectar con el servidor',
            })
          }

          // fetch('http://localhost:3000/api/citas/' + arg.event.id, {
          //   method: 'DELETE',
          //   headers: {
          //     'Content-Type': 'application/json'
          //   }
          // })
          //   .then(response => response.json())
          //   .then(data => {
          //     console.log(data)
          //     if (data.success) {
          //       Swal.fire({
          //         icon: 'success',
          //         title: 'Cita eliminada correctamente',
          //         showConfirmButton: false,
          //         timer: 1500
          //       })
          //       arg.event.remove()
          //     } else {
          //       Swal.fire({
          //         icon: 'error',
          //         title: 'Oops...',
          //         text: 'No se pudo eliminar!',
          //       })
          //     }
          //   })
        }
      })
    },

    events: function (fetchInfo, successCallback, failureCallback) {
      let fechaActual = new Date();
      // fechaDesfasada = fechaDesfasada.getTime();
      fechaActual = fechaActual.getTime();
      console.log(fechaActual);
      // let resta = fechaActual - fechaDesfasada;

      // sumar los milisegundos a la fecha actual
      let fechaDesfasada = new Date(fetchInfo.startStr);
      fechaDesfasada = fechaDesfasada.getTime();
      fechaDesfasada = new Date(fechaDesfasada + + 1610799915);


      let year = fechaDesfasada.getFullYear();
      console.log({ year });

      console.log({ year })
      // Construye la URL de la solicitud
      var url = 'http://localhost:3000/api/citas/year/' + year;

      // Realiza la solicitud HTTP GET
      fetch(url, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        }
      })
        .then(response => response.json())
        .then(data => {
          // Procesa los datos recibidos y llama a successCallback
          var events = data.map(function (item) {
            return {
              title: item.titulo,
              start: item.fecha_inicio,
              end: item.fecha_fin,
              id: item.ID,
            };
          });
          successCallback(events);
        })
        .catch(error => {
          // Maneja los errores y llama a failureCallback si es necesario
          console.error('Error al cargar eventos:', error);
          failureCallback(error);
        });
    },

  });

  calendar.render();
});

function formatDate(fechaOriginal) {

  let separator = fechaOriginal.indexOf('T')
  let date = fechaOriginal.slice(0, separator)
  let hoursMinutesAndSeconds = fechaOriginal.slice(separator + 1, fechaOriginal.length - 5)

  return `${date} ${hoursMinutesAndSeconds}`

}