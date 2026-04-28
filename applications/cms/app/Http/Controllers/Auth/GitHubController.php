<?php

namespace app\Http\Controllers\Auth;

use App\Http\Controllers\Controller;
use App\Models\User;
use Filament\Facades\Filament;
use Illuminate\Support\Facades\Auth;
use Laravel\Socialite\Facades\Socialite;

class GitHubController extends Controller
{
    public function redirectToGitHub()
    {
        return Socialite::driver('github')->redirect();
    }

    public function handleGitHubCallback()
    {
        $githubUser = Socialite::driver('github')->user();

        $user = User::query()
            ->where('email', $githubUser->getEmail())
            ->first();

        if (!$user) {
            return redirect(Filament::getPanel('admin')->getLoginUrl())
                ->withErrors([
                    'email' => 'Nenhum usuário cadastrado com esta conta do GitHub.',
                ]);
        }

        if (!$user->github_id) {
            $user->update([
                'github_id' => $githubUser->getId(),
            ]);
        }

        Auth::login($user);

        $url = Filament::getPanel('admin')->getUrl();

        return redirect()->intended($url);
    }
}
