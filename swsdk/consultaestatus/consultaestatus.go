package consultaestatus

// ConsultarEstatus consulta el estatus de un CFDI en el SAT
func ConsultarEstatus(rfcEmisor, rfcReceptor, total, uuid string) (*ConsultaEstatusResponse, error) {
	return ConsultarEstatusHelper(rfcEmisor, rfcReceptor, total, uuid)
}
