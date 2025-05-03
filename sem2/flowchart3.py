# 1. meminta user input radius
r = float(input("Masukkan radius: "))

# 2. meminta user input tinggi
t = float(input("Masukkan tinggi: "))

# 3. volume tabung dihitung menggunakan rumus:
pi = 3.14 # deklarasi nilai pi
volume_tabung = pi * (r ** 2) * t # rumus volume = pi * r^2 * t

# 4. menampilkan output hasil
print(f"Volume tabung adalah: {volume_tabung:.2f}")
