<?php

namespace Tests\Feature;

use Illuminate\Foundation\Testing\RefreshDatabase;
use Illuminate\Foundation\Testing\WithFaker;
use Tests\TestCase;

class UserTest extends TestCase
{
    /**
     * A basic feature test example.
     *
     * @return voida
     */
//    public function test_example()
//    {
//        $response = $this->get('/');
//
//        $response->assertStatus(200);
//    }

    public function test_createUser()
    {
        $data = [
            'name' => "New User",
            'mobile_number' => "8212233446",
            'email' => "newuser11@gmail.com",
        ];
        //$user = factory(\App\User::class)->create();
        $response = $this->json('POST', '/api/users', $data);
        $response->assertStatus(200);
        $response->assertJson(["status" => true]);
        $response->assertJson(["message" => "User Created!"]);
        $response->assertJson(["data" => $data]);
    }

    public function test_getAllUsers()
    {
        $response = $this->json('GET', '/api/users');
        $response->assertStatus(200);

        $response->assertJsonStructure(
            [
                [
                    "id",
                    "name",
                    "mobile_number",
                    "email",
                ]
            ]
        );
    }

    public function testDeleteUserById()
    {
        $id = rand(1, 10000);
        $response = $this->json('DELETE', "/api/users/$id");
        $response->assertStatus(200);

        $response->assertStatus(200);
        $response->assertJson(['message' => "User with id-$id deleted"]);
    }

    public function testGetUserById()
    {
        $response = $this->json('GET', '/api/users');
        $response->assertStatus(200);
    }
}
