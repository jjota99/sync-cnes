package syncronize

import (
	"main.go/app/core/common/process_csv"
)

func EquipeSyncronize() {
	var headers = []string{
		"id_municipio",
		"id_area",
		"id_origem",
		"id_unidade",
		"id_tipo_equipe",
		"id_sub_tipo_equipe",
		"referencia",
		"dt_ativacao",
		"dt_desativacao",
		"tp_pop_assist_quilomb",
		"tp_pop_assist_assent",
		"tp_pop_assist_geral",
		"tp_pop_assist_escola",
		"tp_pop_assist_pronasci",
		"tp_pop_assist_indigena",
		"tp_pop_assist_ribeirinha",
		"tp_pop_assist_situacao_rua",
		"tp_pop_assist_priv_liberdade",
		"tp_pop_assist_conflito_lei",
		"tp_pop_assist_adol_conf_lei",
		"id_cnes_uom",
		"nu_ch_amb_uom",
		"id_motivo_desativ",
		"id_tipo_desativ",
		"ine",
		"dt_atualizacao",
		"usuario",
		"dt_atualizacao_origem",
	}

	var excludeColumns = []string{"CO_PROF_SUS_PRECEPTOR", "CO_CNES_PRECEPTOR"}

	process_csv.ProcessCsv("tbEquipe", headers, excludeColumns)
}
