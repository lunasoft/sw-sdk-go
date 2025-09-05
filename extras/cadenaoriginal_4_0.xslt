<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="2.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:cfdi="http://www.sat.gob.mx/cfd/4">
<xsl:output method="text" version="1.0" encoding="UTF-8" indent="no"/>
<xsl:template match="/|@*|node()">
<xsl:apply-templates select="@*|node()"/>
</xsl:template>
<xsl:template match="text()"/>
<xsl:template match="cfdi:Comprobante">
<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Version"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Serie"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Folio"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Fecha"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@FormaPago"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@NoCertificado"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@CondicionesDePago"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@SubTotal"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Moneda"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Total"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@TipoDeComprobante"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@MetodoPago"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@LugarExpedicion"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Exportacion"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Confirmacion"/>
</xsl:call-template>|<xsl:apply-templates select="./cfdi:Emisor"/>|<xsl:apply-templates select="./cfdi:Receptor"/>|<xsl:apply-templates select="./cfdi:Conceptos"/>|<xsl:apply-templates select="./cfdi:Impuestos"/>|<xsl:apply-templates select="./cfdi:Complemento"/>
</xsl:template>
<xsl:template match="cfdi:Emisor">
<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Rfc"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Nombre"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@RegimenFiscal"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@FacAtrAdquirente"/>
</xsl:call-template>
</xsl:template>
<xsl:template match="cfdi:Receptor">
<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Rfc"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Nombre"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@DomicilioFiscalReceptor"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@RegimenFiscalReceptor"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@UsoCFDI"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@DomicilioFiscalReceptor"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@RegimenFiscalReceptor"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@UsoCFDI"/>
</xsl:call-template>
</xsl:template>
<xsl:template match="cfdi:Conceptos">
<xsl:apply-templates select="./cfdi:Concepto"/>
</xsl:template>
<xsl:template match="cfdi:Concepto">
<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@ClaveProdServ"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Cantidad"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@ClaveUnidad"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Descripcion"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@ValorUnitario"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Importe"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Descuento"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@ObjetoImp"/>
</xsl:call-template>|<xsl:apply-templates select="./cfdi:Impuestos"/>
</xsl:template>
<xsl:template match="cfdi:Impuestos">
<xsl:apply-templates select="./cfdi:Retenciones"/>
<xsl:apply-templates select="./cfdi:Traslados"/>
</xsl:template>
<xsl:template match="cfdi:Retenciones">
<xsl:apply-templates select="./cfdi:Retencion"/>
</xsl:template>
<xsl:template match="cfdi:Retencion">
<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Base"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Impuesto"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@TipoFactor"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@TasaOCuota"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Importe"/>
</xsl:call-template>
</xsl:template>
<xsl:template match="cfdi:Traslados">
<xsl:apply-templates select="./cfdi:Traslado"/>
</xsl:template>
<xsl:template match="cfdi:Traslado">
<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Base"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Impuesto"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@TipoFactor"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@TasaOCuota"/>
</xsl:call-template>|<xsl:call-template name="Requerido">
<xsl:with-param name="valor" select="./@Importe"/>
</xsl:call-template>
</xsl:template>
<xsl:template match="cfdi:Impuestos">
<xsl:apply-templates select="./cfdi:Retenciones"/>
<xsl:apply-templates select="./cfdi:Traslados"/>
</xsl:template>
<xsl:template match="cfdi:Complemento">
<xsl:apply-templates select="./*"/>
</xsl:template>
<xsl:template name="Requerido">
<xsl:param name="valor"/>
<xsl:choose>
<xsl:when test="$valor">
<xsl:value-of select="$valor"/>
</xsl:when>
<xsl:otherwise>
</xsl:otherwise>
</xsl:choose>
</xsl:template>
</xsl:stylesheet>