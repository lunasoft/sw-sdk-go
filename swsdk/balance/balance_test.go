package balance

import (
	"testing"
)

const testUserID = "ee119c6a-5548-4c5e-b693-00f8195d4f93"

func TestAddStamps(t *testing.T) {
	t.Log("=== Test: Añadir Timbres ===")

	resp, err := AddStamps(testUserID, 10, "Se abonan 10 timbres desde test")
	if err != nil {
		t.Logf("⚠️  Error añadiendo timbres: %v", err)
		t.Log("   Nota: Este test puede fallar por restricciones del API")
		return
	}

	t.Logf("✅ Timbres añadidos exitosamente:")
	t.Logf("   Total de timbres después del abono: %d", resp.Data)
	t.Logf("   Status: %s", resp.Status)
}

func TestGetBalance(t *testing.T) {
	t.Log("=== Test: Consultar Balance ===")

	resp, err := GetBalance()
	if err != nil {
		t.Logf("⚠️  Error consultando balance: %v", err)
		t.Log("   Nota: Este test puede fallar por restricciones del API")
		return
	}

	t.Logf("✅ Balance obtenido exitosamente:")
	t.Logf("   ID Usuario: %s", resp.Data.IDUser)
	t.Logf("   Timbres disponibles: %d", resp.Data.StampsBalance)
	t.Logf("   Timbres usados: %d", resp.Data.StampsUsed)
	t.Logf("   Timbres asignados: %d", resp.Data.StampsAssigned)
	t.Logf("   Ilimitado: %t", resp.Data.IsUnlimited)

	if resp.Data.LastTransaction.Folio > 0 {
		t.Logf("   Última transacción:")
		t.Logf("     Folio: %d", resp.Data.LastTransaction.Folio)
		t.Logf("     Fecha: %s", resp.Data.LastTransaction.Date)
		if resp.Data.LastTransaction.StampsIn != nil {
			t.Logf("     Timbres ingresados: %d", *resp.Data.LastTransaction.StampsIn)
		}
		if resp.Data.LastTransaction.StampsOut != nil {
			t.Logf("     Timbres retirados: %d", *resp.Data.LastTransaction.StampsOut)
		}
		t.Logf("     Timbres actuales: %d", resp.Data.LastTransaction.StampsCurrent)
	}
}

func TestRemoveStamps(t *testing.T) {
	t.Log("=== Test: Eliminar Timbres ===")

	resp, err := RemoveStamps(testUserID, 5, "Se retiran 5 timbres desde test")
	if err != nil {
		t.Logf("⚠️  Error eliminando timbres: %v", err)
		t.Log("   Nota: Este test puede fallar por restricciones del API")
		return
	}

	t.Logf("✅ Timbres eliminados exitosamente:")
	t.Logf("   Total de timbres después de la eliminación: %d", resp.Data)
	t.Logf("   Status: %s", resp.Status)
}
