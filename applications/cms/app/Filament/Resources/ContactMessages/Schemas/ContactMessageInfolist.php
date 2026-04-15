<?php

namespace App\Filament\Resources\ContactMessages\Schemas;

use Filament\Infolists\Components\TextEntry;
use Filament\Schemas\Components\Section;
use Filament\Schemas\Schema;

class ContactMessageInfolist
{
    public static function configure(Schema $schema): Schema
    {
        return $schema
            ->components([
                Section::make('Message Details')
                    ->schema([
                        TextEntry::make('name'),
                        TextEntry::make('email')
                            ->copyable(),
                        TextEntry::make('status')
                            ->badge(),
                        TextEntry::make('created_at')
                            ->dateTime(),
                    ])->columns(2),

                Section::make('Conteúdo da Mensagem')
                    ->schema([
                        TextEntry::make('message')
                            ->label('')
                            ->prose(),
                    ]),
            ]);
    }
}
