package main

import "fmt"

func convertWaktuMars(detikBumi int) (int, int, int) {
	const jamSatuanDetikMars = 4500
	const jamSatuanMenitMars = 75

	jamMars := detikBumi / jamSatuanDetikMars
	sisaDetik := detikBumi % jamSatuanDetikMars
	menitMars := sisaDetik/jamSatuanMenitMars
	sisaDetik = sisaDetik % jamSatuanDetikMars
	detikMars := sisaDetik

return jamMars, menitMars, detikMars
}

func main() {
	var detikBumi int
	fmt.Print("Masukkan detik: ")
	fmt.Scan(&detikBumi)

	jamMars, menitMars, detikMars := convertWaktuMars(detikBumi)

	fmt.Printf("Hasil konversi %d jam, %d menit, %d detik di Mars\n", jamMars, menitMars, detikMars)
}


// nnyalin aja
onst canvas = document.querySelector('canvas')
  const c = canvas.getContext('2d')

canvas.width = 1024
canvas.height= 576

class Sprite {
  constructor({position}) {
    this.position = position
    this.image= new Image()
    this.image.src = './img/qwerty1map.png'
  }
  draw() {
    c.drawImage(this.image, this.position.x, this.position.y)
  }
}

const level1cvmap = new Sprite({
  position: {
    x: 0,
    y: 0,
  },
})

const player = new Player()

const keys = {
  w: {
    pressed: false
  },
  a: {
    pressed: false
  },
  d: {
  pressed: false
  }
}
  function animate () {
    window.requestAnimationFrame(animate)
    c.fillStyle = 'white'
    c.fillRect(0, 0, canvas.width, canvas.height)

    qwerty1map.png.draw()

    player.velocity.x = 0
if (keys.d.pressed) player.velocity.x = 4
 else if (keys.a.pressed) player.velocity.x = -4

player.draw()
player.update()
}

animate()

// ini player.js
class Player {
  constructor() {
    this.position = {
      x: 100,
      y: 100,
    }

    this.velocity = {
      x: 0,
      y: 0,
    }
    class Sprite {
      constructor({position}) {
        this.position = position
        this.image = new Image()
        this.image.src = './img/RoniNYoba.png'
        this.loaded = false
    
    this.width = 100
    this.height = 100   
    this.sides = {
      bottom: this.position.y + this.height
    }
    this.gravity = 1
  }

  draw() {
    c.fillStyle = 'red'
    c.fillRect(this.position.x, this.position.y, this.width, this.height)
  }

  update() {
    this.position.x += this.velocity.x
    this.position.y += this.velocity.y
    this.sides.bottom = this.position.y + this.height

    //above bottom of canvas
    if (this.sides.bottom +this.velocity.y< canvas.height) {
      this.velocity.y += this.gravity
    }  else this.velocity.y = 0
  }
} 