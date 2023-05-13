<?php

namespace App\Http\Controllers;

use Illuminate\Support\Facades\Http;

class MerchantController extends Controller
{
    // asumsi kita ingin get data merchant dengan shop id 01
    public function getDetailMerchant() {
        $response = Http::get("localhost:9000/v1/shops/01");
        $data = $response->json();

        $merchantDetail = [
            "id" => $data["id"],
            "owner" => "Andromeda",
            "address" => "Bogor",
            "shop_name" => $data["name"],
            "number_of_product" => $data["number_of_product"]
        ];

        return response()->json($merchantDetail, 200);
    }
}
