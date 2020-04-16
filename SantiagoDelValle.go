package main

func main() {

	//NOTA: CUANDO SE REFIERE A PROMEDIO SE REFIERE A PORCENTAJE

	suma := 27834 + 23789 + 30189 + 30967 + 32501 + 32701 + 31665 + 17155 + 4614 + 834

	promedio0 := (27834 * 100 / suma)
	promedio1 := (23789 * 100 / suma)
	promedio2 := (30189 * 100 / suma)
	promedio3 := (30967 * 100 / suma)
	promedio4 := (32501 * 100 / suma)
	promedio5 := (32701 * 100 / suma)
	promedio6 := (31665 * 100 / suma)
	promedio7 := (17155 * 100 / suma)
	promedio8 := (4614 * 100 / suma)
	promedio9 := (834 * 100 / suma)

	PromAcumulado := promedio0 + promedio1 + promedio2 + promedio3 + promedio4 + promedio5 + promedio6 + promedio7 + promedio8 + promedio9

	println("El porcentaje del año 2008 con respecto al total fue de", promedio0)
	println("El porcentaje del año 2009 con respecto al total fue de", promedio1)
	println("El porcentaje del año 2010 con respecto al total fue de", promedio2)
	println("El porcentaje del año 2011 con respecto al total fue de", promedio3)
	println("El porcentaje del año 2012 con respecto al total fue de", promedio4)
	println("El porcentaje del año 2013 con respecto al total fue de", promedio5)
	println("El porcentaje del año 2014 con respecto al total fue de", promedio6)
	println("El porcentaje del año 2015 con respecto al total fue de", promedio7)
	println("El porcentaje del año 2016 con respecto al total fue de", promedio8)
	println("El porcentaje del año 2017 con respecto al total fue de", promedio9)
	println("El porcentaje acumulado fue de..", PromAcumulado)
}
