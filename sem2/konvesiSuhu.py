def fahrenheitToCelsius(f): # konversi fahrenheit-celsius
    return (5/9) * (f - 32)

def kelvinToReamur(k): # konversi kelvin-reamur
    return (4/5) * (k - 273)

# nilai suhu + nim
fahrenheit = 87  # 50 + 37
kelvin = 512  # 475 + 37

# call function
hasilCelsius = fahrenheitToCelsius(fahrenheit)
hasilReamur = kelvinToReamur(kelvin)

# menampilkan output hasil konversi
print(f"{fahrenheit}°Fahrenheit = {hasilCelsius:.2f}°Celsius")
print(f"{kelvin}°Kelvin = {hasilReamur:.2f}°Reamur")
